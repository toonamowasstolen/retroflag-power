package logging

import (
	"fmt"
	"io"
	"log"
)

// New returns the daemon's standard logger, writing to w.
func New(w io.Writer) *log.Logger {
	return log.New(w, "", log.LstdFlags)
}

const bannerRule = "----------------------------------------"

// Banner returns the one-time startup banner for appName/version, e.g.
// "retroflag-powerd 0.1.0-dev". Plain ASCII, no color, no box-drawing —
// matches PROJECT_MANIFEST.md's "Terminal Personality" guidance (warm,
// occasionally playful, not noisy) and PROJECT_MEMORY.md's "Good" tone
// example. Printed once at real daemon startup, never on the diagnostic/
// dry-run/fake-event command paths, which have their own parseable,
// banner-free stdout contracts.
func Banner(appName, version string) string {
	return fmt.Sprintf("%s\n%s %s\nPlayer One, get ready.\n%s\n", bannerRule, appName, version, bannerRule)
}

// WriteBanner writes Banner's output to w.
func WriteBanner(w io.Writer, appName, version string) {
	fmt.Fprint(w, Banner(appName, version))
}
