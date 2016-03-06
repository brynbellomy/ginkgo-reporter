

# gomega-reporter

A test reporter for the Gomega testing framework.

Use the following lines to initialize your test suite:

```go
package blah_test

import (
    "github.com/brynbellomy/gomega-reporter"
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"

    "testing"
)

func TestBlah(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecsWithCustomReporters(t, "Blah Suite", []Reporter{
        &reporter.TerseReporter{Logger: &reporter.DefaultLogger{}},
    })
}

```

# license

ISC