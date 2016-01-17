package builder

import (
	"os/exec"
	"strings"
)

func getPackageName(path string) string {
	output, err := exec.Command("go", "list", "-f", "{{.Name}}", path).CombinedOutput()
	if err != nil {
		return guessPackageName(path)
	}
	return strings.TrimSpace(string(output))
}

func guessPackageName(path string) string {
	preferred := path
	if strings.HasSuffix(preferred, "/") {
		// training slashes are usually tolerated, so we can get rid of one if it exists
		preferred = preferred[:len(preferred)-1]
	}
	if strings.Contains(preferred, "/") {
		// if the path contains a "/", use the last part
		preferred = preferred[strings.LastIndex(preferred, "/")+1:]
	}
	if strings.Contains(preferred, "-") {
		// the name usually follows a hyphen - e.g. github.com/foo/go-bar
		// if the package name contains a "-", use the last part
		preferred = preferred[strings.LastIndex(preferred, "-")+1:]
	}
	if strings.Contains(preferred, ".") {
		// dot is commonly usually used as a version - e.g. github.com/foo/bar.v1
		// if the package name contains a ".", use the first part
		preferred = preferred[:strings.LastIndex(preferred, ".")]
	}
	return preferred
}
