package main

import "tholian-core/matchers"
import "tholian-core/structs"
import "tholian-scout/actions"
import "tholian-scout/adapters/system/proc"
import "encoding/json"
import "fmt"
import "os"
import "path"
import "strings"

var CONSOLE *structs.Console
var SYSTEM *structs.System = nil

func showUsage(console *structs.Console) {

	console.Info("")
	console.Info("Tholian Scout")
	console.Info("Software, Hardware, Network and User Inventory Agent")
	console.Info("")

	console.Group("Usage: tholian-scout [Action]")
	console.Log("")
	console.Log("The \"search\" Action can be used to identify affected objects on the system.")
	console.Log("It supports the \"antiques\", \"devices\", \"incidents\", \"networks\",")
	console.Log("\"packages\", \"programs\", \"services\", \"updates\", \"users\" types.")
	console.Log("")
	console.GroupEnd("-------")

	console.Group("Action  | Description                              |")
	console.Log("--------|------------------------------------------|")
	console.Log("install | Installs Tholian Scout on the system.    |")
	console.Log("collect | Collects a System Report.                |")
	console.Log("search  | Searches the System Report for patterns. |")
	console.GroupEnd("--------|------------------------------------------|")

	console.Group("Examples")
	console.Log("# Collect System Report")
	console.Log("sudo tholian-scout collect;")
	console.Log("sudo tholian-scout collect /tmp/report.json;")
	console.Log("")
	console.Log("# Search for vulnerable log4j versions")
	console.Log("tholian-scout search antiques log4j;")
	console.Log("tholian-scout search antiques \"log4j <= 2.14.1\";")
	console.Log("")
	console.Log("# Search for glibc packages")
	console.Log("tholian-scout search packages glibc;")
	console.Log("tholian-scout search packages \"glibc = 2.36\";")
	console.GroupEnd("--------")

}

func init() {
	CONSOLE = structs.NewConsole(os.Stdout, os.Stderr, 0)
	SYSTEM = actions.Init(CONSOLE)
}

func isSearchType(value string) bool {

	var result bool = false

	types := []string{
		"antiques",
		"devices",
		"drives",
		"incidents",
		"networks",
		"packages",
		"programs",
		"services",
		"updates",
		"users",
	}

	for t := 0; t < len(types); t++ {

		if value == types[t] {
			result = true
			break
		}

	}

	return result

}

func isReportPath(chunk string) bool {

	var result bool = false

	cwd, err := os.Getwd()

	if err == nil {

		resolved := path.Join(cwd, chunk)

		if strings.HasPrefix(resolved, "/home") || strings.HasPrefix(resolved, "/tmp") || strings.HasPrefix(resolved, "/var") {

			if strings.HasSuffix(resolved, ".json") {
				result = true
			}

		}

	}

	return result

}

func main() {

	console := CONSOLE

	var action string
	var search_type string
	var search_data string

	if len(os.Args) == 2 {

		action = strings.ToLower(os.Args[1])

	} else if len(os.Args) == 3 && os.Args[1] == "collect" {

		action = "collect"

		if isReportPath(os.Args[2]) {

			// "tholian-scout collect /tmp/report.json"
			os.Setenv("THOLIAN_REPORT", os.Args[2])

		} else {

			check := os.Getenv("THOLIAN_REPORT")

			if isReportPath(check) {
				// Do Nothing
			} else {
				os.Setenv("THOLIAN_REPORT", "")
			}

		}

	} else if len(os.Args) == 4 && os.Args[1] == "search" {

		action = "search"

		if isSearchType(os.Args[2]) {
			// "tholian-scout search antiques log4j"
			search_type = os.Args[2]
			search_data = strings.TrimSpace(os.Args[3])
		}

	}

	console.ClearScreen()

	console.Group("tholian-scout: Command-Line Arguments")

	if action == "search" {
		console.Inspect(struct {
			Action     string
			SearchType string
			SearchData string
		}{
			Action:     action,
			SearchType: search_type,
			SearchData: search_data,
		})
	} else {
		console.Inspect(struct{ Action string }{Action: action})
	}

	console.GroupEnd("")

	if proc.Init(console) == true {

		if action == "install" {

			result := actions.Install(CONSOLE, SYSTEM)

			if result == true {
				os.Exit(0)
			} else {
				os.Exit(1)
			}

		} else if action == "uninstall" {

			result := actions.Uninstall(CONSOLE, SYSTEM)

			if result == true {
				os.Exit(0)
			} else {
				os.Exit(1)
			}

		} else if action == "collect" {

			result := actions.Collect(CONSOLE, SYSTEM)

			if result == true {

				report := os.Getenv("THOLIAN_REPORT")

				if report != "" {

					data, err0 := json.MarshalIndent(SYSTEM, "", "\t")

					if err0 == nil {

						err1 := os.WriteFile(report, data, 0644)

						if err1 == nil {
							console.Info(fmt.Sprintf("Report written to %s", report))
							os.Exit(0)
						} else {
							console.Error(err1.Error())
							os.Exit(1)
						}

					} else {
						console.Error(err0.Error())
						os.Exit(1)
					}

				} else {

					data, err0 := json.MarshalIndent(SYSTEM, "", "\t")

					if err0 == nil {

						console.Inspect(data)
						os.Exit(0)

					} else {
						console.Error(err0.Error())
						os.Exit(1)
					}

				}

			} else {
				os.Exit(1)
			}

		} else if action == "search" && search_type != "" && search_data != "" {

			if search_type == "antiques" {

				actions.CollectServices(CONSOLE, SYSTEM)
				actions.CollectPackages(CONSOLE, SYSTEM)
				actions.CollectUpdates(CONSOLE, SYSTEM)
				actions.CollectAntiques(CONSOLE, SYSTEM)

				results := actions.SearchAntiques(CONSOLE, SYSTEM, matchers.ToAntique(search_data))

				if len(results) > 0 {

					for r := 0; r < len(results); r++ {
						console.Inspect(results[r])
					}

					os.Exit(0)

				} else {
					os.Exit(1)
				}

			} else if search_type == "devices" {

				actions.CollectDevices(CONSOLE, SYSTEM)

				results := actions.SearchDevices(CONSOLE, SYSTEM, matchers.ToDevice(search_data))

				if len(results) > 0 {

					for r := 0; r < len(results); r++ {
						console.Inspect(results[r])
					}

					os.Exit(0)

				} else {
					os.Exit(1)
				}

			} else if search_type == "drives" {

				actions.CollectDrives(CONSOLE, SYSTEM)

				results := actions.SearchDrives(CONSOLE, SYSTEM, matchers.ToDrive(search_data))

				if len(results) > 0 {

					for r := 0; r < len(results); r++ {
						console.Inspect(results[r])
					}

					os.Exit(0)

				} else {
					os.Exit(1)
				}

			} else if search_type == "incidents" {

				actions.CollectIncidents(CONSOLE, SYSTEM)

				results := actions.SearchIncidents(CONSOLE, SYSTEM, matchers.ToIncident(search_data))

				if len(results) > 0 {

					for r := 0; r < len(results); r++ {
						console.Inspect(results[r])
					}

					os.Exit(0)

				} else {
					os.Exit(1)
				}

			} else if search_type == "networks" {

				actions.CollectNetworks(CONSOLE, SYSTEM)

				results := actions.SearchNetworks(CONSOLE, SYSTEM, matchers.ToNetwork(search_data))

				if len(results) > 0 {

					for r := 0; r < len(results); r++ {
						console.Inspect(results[r])
					}

					os.Exit(0)

				} else {
					os.Exit(1)
				}

			} else if search_type == "packages" {

				actions.CollectPackages(CONSOLE, SYSTEM)

				results := actions.SearchPackages(CONSOLE, SYSTEM, matchers.ToPackage(search_data))

				if len(results) > 0 {

					for r := 0; r < len(results); r++ {
						console.Inspect(results[r])
					}

					os.Exit(0)

				} else {
					os.Exit(1)
				}

			} else if search_type == "programs" {

				actions.CollectPrograms(CONSOLE, SYSTEM)

				results := actions.SearchPrograms(CONSOLE, SYSTEM, matchers.ToProgram(search_data))

				if len(results) > 0 {

					for r := 0; r < len(results); r++ {
						console.Inspect(results[r])
					}

					os.Exit(0)

				} else {
					os.Exit(1)
				}

			} else if search_type == "services" {

				actions.CollectPrograms(CONSOLE, SYSTEM)
				actions.CollectPackages(CONSOLE, SYSTEM)
				actions.CollectUsers(CONSOLE, SYSTEM)

				results := actions.SearchServices(CONSOLE, SYSTEM, matchers.ToProgram(search_data))

				if len(results) > 0 {

					for r := 0; r < len(results); r++ {
						console.Inspect(results[r])
					}

					os.Exit(0)

				} else {
					os.Exit(1)
				}

			} else if search_type == "updates" {

				actions.CollectUpdates(CONSOLE, SYSTEM)

				results := actions.SearchUpdates(CONSOLE, SYSTEM, matchers.ToUpdate(search_data))

				if len(results) > 0 {

					for r := 0; r < len(results); r++ {
						console.Inspect(results[r])
					}

					os.Exit(0)

				} else {
					os.Exit(1)
				}

			} else if search_type == "users" {

				actions.CollectUsers(CONSOLE, SYSTEM)

				results := actions.SearchUsers(CONSOLE, SYSTEM, matchers.ToUser(search_data))

				if len(results) > 0 {

					for r := 0; r < len(results); r++ {
						console.Inspect(results[r])
					}

					os.Exit(0)

				} else {
					os.Exit(1)
				}

			} else {

				showUsage(CONSOLE)

				os.Exit(1)

			}

		} else {

			CONSOLE.ClearScreen()
			showUsage(CONSOLE)

			os.Exit(1)

		}

	} else {

		showUsage(CONSOLE)

		os.Exit(1)

	}

}
