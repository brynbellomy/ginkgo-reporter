package reporter

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/types"
)

type TerseReporter struct {
	Logger       ILogger
	DisableColor bool
}

type ILogger interface {
	Infoln(args ...interface{})
	Errorln(args ...interface{})
}

func (r *TerseReporter) SpecSuiteWillBegin(config config.GinkgoConfigType, s *types.SuiteSummary) {
	msg := color.WhiteString("<" + s.SuiteDescription + ">")
	r.Logger.Infoln(msg)
}

func (r *TerseReporter) BeforeSuiteDidRun(setupSummary *types.SetupSummary) {
	// no-op
}

func (r *TerseReporter) SpecWillRun(specSummary *types.SpecSummary) {
	// no-op
}

func (r *TerseReporter) SpecDidComplete(specSummary *types.SpecSummary) {
	var prefix string
	var c, bold *color.Color

	msg := strings.Join(specSummary.ComponentTexts[1:], ", ")
	if specSummary.HasFailureState() {
		prefix = "    (fail)"
		c = color.New(color.FgRed)
		bold = color.New(color.FgRed, color.Bold)

	} else if specSummary.State == types.SpecStatePending {
		prefix = "    (pend)"
		c = color.New(color.FgYellow)
		bold = color.New(color.FgYellow, color.Bold)

	} else if specSummary.State == types.SpecStateSkipped {
		prefix = "    (skip)"
		c = color.New(color.FgBlue)
		bold = color.New(color.FgBlue, color.Bold)

	} else {
		prefix = "    (pass)"
		c = color.New(color.FgGreen)
		bold = color.New(color.FgGreen, color.Bold)
	}

	if r.DisableColor {
		c.DisableColor()
		bold.DisableColor()
	}

	r.Logger.Infoln(bold.SprintFunc()(prefix) + " " + c.SprintFunc()(msg) + " " + fmt.Sprintf("%v", specSummary.RunTime))
	if specSummary.HasFailureState() {
		red := color.New(color.FgRed)
		r.Logger.Errorln(red.SprintFunc()(specSummary.Failure.Message))
		r.Logger.Errorln(red.SprintFunc()(specSummary.Failure.ForwardedPanic))
		r.Logger.Errorln(red.SprintfFunc()("%v:%v", specSummary.Failure.Location.FileName, specSummary.Failure.Location.LineNumber))
		r.Logger.Errorln(red.SprintFunc()(specSummary.Failure.Location.FullStackTrace))
	}
}

func (r *TerseReporter) AfterSuiteDidRun(setupSummary *types.SetupSummary) {
	// no-op
}

func (r *TerseReporter) SpecSuiteDidEnd(s *types.SuiteSummary) {
	msg := color.WhiteString("</" + s.SuiteDescription + ">")
	r.Logger.Infoln(msg)

	strs := []string{}
	strs = append(strs, fmt.Sprintf("%v %v", s.NumberOfPassedSpecs, color.GreenString("passed")))
	strs = append(strs, fmt.Sprintf("%v %v", s.NumberOfFailedSpecs, color.RedString("failed")))
	strs = append(strs, fmt.Sprintf("%v %v", s.NumberOfPendingSpecs, color.YellowString("pending")))
	strs = append(strs, fmt.Sprintf("%v %v", s.NumberOfSkippedSpecs, color.BlueString("skipped")))
	strs = append(strs, fmt.Sprintf("%v %v", s.NumberOfTotalSpecs, color.WhiteString("total")))
	r.Logger.Infoln(strings.Join(strs, ", "))
	boldw := color.New(color.FgHiWhite)
	r.Logger.Infoln(fmt.Sprintf("Total run time: %v", boldw.SprintFunc()(s.RunTime)))
}
