package qemu

import "crypto/sha256"
import "encoding/hex"
import "fmt"
import "io"
import "net/http"
import "os"
import "os/exec"
import "path/filepath"
import "time"

const ImageURL = "https://download.freebsd.org/releases/VM-IMAGES/15.1-RELEASE/amd64/Latest/FreeBSD-15.1-RELEASE-amd64-BASIC-CLOUDINIT-ufs.qcow2.xz"
const ImageSHA256 = "e4ca4db889f8559c9b9dfcacc70405c038476f4b6d41649b152d3809a2ed9e1f"
const ImageName = "FreeBSD-15.1-RELEASE-amd64-BASIC-CLOUDINIT-ufs.qcow2"

func EnsureImage(cacheDir string) (string, error) {

	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		return "", err
	}

	image := filepath.Join(cacheDir, ImageName)

	if stat, err := os.Stat(image); err == nil && stat.Size() > 0 {
		return image, nil
	}

	archive := image + ".xz"

	if stat, err := os.Stat(archive); err != nil || stat.Size() == 0 {
		if err := download(ImageURL, archive); err != nil {
			return "", err
		}
	}

	if err := verifySHA256(archive, ImageSHA256); err != nil {
		return "", err
	}

	cmd := exec.Command("xz", "-d", "-k", "-f", archive)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("decompress image: %w", err)
	}

	return image, nil

}

func download(url string, destination string) error {

	client := &http.Client{Timeout: 30 * time.Minute}

	response, err := client.Get(url)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("download image: %s", response.Status)
	}

	partial := destination + ".part"
	file, err := os.Create(partial)

	if err != nil {
		return err
	}

	if _, err := io.Copy(file, response.Body); err != nil {
		file.Close()
		os.Remove(partial)
		return err
	}

	if err := file.Close(); err != nil {
		os.Remove(partial)
		return err
	}

	return os.Rename(partial, destination)

}

func verifySHA256(name string, expected string) error {

	file, err := os.Open(name)

	if err != nil {
		return err
	}

	defer file.Close()

	hash := sha256.New()

	if _, err := io.Copy(hash, file); err != nil {
		return err
	}

	actual := hex.EncodeToString(hash.Sum(nil))

	if actual != expected {
		return fmt.Errorf("checksum mismatch for %s: expected %s, got %s", name, expected, actual)
	}

	return nil

}

func CreateOverlay(base string, overlay string) error {

	absolute, err := filepath.Abs(base)

	if err != nil {
		return err
	}

	cmd := exec.Command("qemu-img", "create", "-f", "qcow2", "-F", "qcow2", "-b", absolute, overlay)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()

}
