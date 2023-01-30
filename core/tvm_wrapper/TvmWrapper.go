
package tvm_wrapper

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"./gotvm"
)

// TvmConfig : TVM Module Configuration
type TvmConfig struct {
	DeviceType int64
}

// NewTvmConfig : Create Object of TvmConfig
func NewTvmConfig() *TvmConfig {
	config := TvmConfig{}
	config.DeviceType = (int64)(gotvm.KDLCPU)
	return &config
}

type moduleInfo struct {
	graphModule *gotvm.Module
	inputShape []int64
	outputShape []int64
}

func newModuleInfo(graphModule *gotvm.Module, inputShape []int64, outputShape []int64) *moduleInfo {
	info := moduleInfo{}
	info.graphModule = graphModule
	info.inputShape = inputShape
	info.outputShape = outputShape
	return &info
}

// ModelParam : model parameters struct to build tvm graph
type ModelParam struct {
	ModelLibPath    string
	ModelJSONPath   string
	ModelParamsPath string
	InputShape      []int64
	OutputShape     []int64
}

// NewModelParam : create instance of model param
func NewModelParam(modelLibPath string, modJsonPath string, modParamsPath string,
	inputShape []int64, outputShape []int64) *ModelParam {
	params := ModelParam{
		modelLibPath,
		modJsonPath,
		modParamsPath,
		inputShape,
		outputShape,
	}
	return &params
}

// DebugStr : get debuggable text of model parameters
func (param *ModelParam) DebugStr() string {
	debugStr := fmt.Sprintf("ModelLibPath : %s\n", param.ModelLibPath)
	debugStr += fmt.Sprintf("ModelJSONPath : %s\n", param.ModelJSONPath)
	debugStr += fmt.Sprintf("ModelParamsPath : %s\n", param.ModelJSONPath)
	debugStr += fmt.Sprintf("InputShape : %v\n", param.InputShape)
	debugStr += fmt.Sprintf("OutputShape : %v\n", param.OutputShape)
	return debugStr
}

// TvmWrapper : TVM wrapper struct to use TVM function
type TvmWrapper struct {
	funcNames []string
	config TvmConfig
}

// NewTvmWrapper : Create TVM wrapper object
func NewTvmWrapper() *TvmWrapper {
	wrapper := TvmWrapper{}
	return &wrapper
}

// Initialize : Initialize TVM wrapper struct to use it
func (wrapper *TvmWrapper) Initialize(config TvmConfig) error {
	defer runtime.GC()

	// display gotvm information
	fmt.Printf("TVM Version   : v%v\n", gotvm.TVMVersion)
	fmt.Printf("DLPACK Version: v%v\n\n", gotvm.DLPackVersion)

	// set configuration
	wrapper.config = config

	// get global function names
	funcNames, err := gotvm.FuncListGlobalNames()
	if err != nil {
		fmt.Print(err.Error())
		return err
	}
	wrapper.funcNames = funcNames
	return nil
}

// LoadModel : Load specified model to get inference model
func (wrapper *TvmWrapper) LoadModel(modelParam *ModelParam) (*moduleInfo, error) {
	defer runtime.GC()

	// debug model parameters