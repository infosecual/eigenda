package deploy

import (
	"testing"

	"github.com/ory/dockertest/v3"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
)

func GetTypeProvider(data []byte) (*go_fuzz_utils.TypeProvider, error) {
	tp, err := go_fuzz_utils.NewTypeProvider(data)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsStringBounds(0, 1024)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsSliceBounds(0, 4096)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsBiases(0, 0, 0, 0)
	if err != nil {
		return nil, err
	}
	return tp, nil
}

func Fuzz_Nosy_Config_DeployExperiment__(f *testing.F) {
	f.Fuzz(func(t *testing.T, testName string, rootPath string) {
		env := NewTestConfig(testName, rootPath)
		env.DeployExperiment()
	})
}

func Fuzz_Nosy_Config_GenerateAllVariables__(f *testing.F) {
	f.Fuzz(func(t *testing.T, testName string, rootPath string) {
		env := NewTestConfig(testName, rootPath)
		env.GenerateAllVariables()
	})
}

func Fuzz_Nosy_Config_GetDeployer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, testName string, rootPath string, name string) {
		env := NewTestConfig(testName, rootPath)
		env.GetDeployer(name)
	})
}

func Fuzz_Nosy_Config_RunNodePluginBinary__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var testName string
		fill_err = tp.Fill(&testName)
		if fill_err != nil {
			return
		}
		var rootPath string
		fill_err = tp.Fill(&rootPath)
		if fill_err != nil {
			return
		}
		var operation string
		fill_err = tp.Fill(&operation)
		if fill_err != nil {
			return
		}
		var operator OperatorVars
		fill_err = tp.Fill(&operator)
		if fill_err != nil {
			return
		}

		env := NewTestConfig(testName, rootPath)
		env.RunNodePluginBinary(operation, operator)
	})
}

func Fuzz_Nosy_Config_SaveTestConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, testName string, rootPath string) {
		env := NewTestConfig(testName, rootPath)
		env.SaveTestConfig()
	})
}

func Fuzz_Nosy_Config_StartAnvil__(f *testing.F) {
	f.Fuzz(func(t *testing.T, testName string, rootPath string) {
		env := NewTestConfig(testName, rootPath)
		env.StartAnvil()
	})
}

func Fuzz_Nosy_Config_StartBinaries__(f *testing.F) {
	f.Fuzz(func(t *testing.T, testName string, rootPath string) {
		env := NewTestConfig(testName, rootPath)
		env.StartBinaries()
	})
}

func Fuzz_Nosy_Config_StartGraphNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, testName string, rootPath string) {
		env := NewTestConfig(testName, rootPath)
		env.StartGraphNode()
	})
}

func Fuzz_Nosy_Config_StopAnvil__(f *testing.F) {
	f.Fuzz(func(t *testing.T, testName string, rootPath string) {
		env := NewTestConfig(testName, rootPath)
		env.StopAnvil()
	})
}

func Fuzz_Nosy_Config_StopBinaries__(f *testing.F) {
	f.Fuzz(func(t *testing.T, testName string, rootPath string) {
		env := NewTestConfig(testName, rootPath)
		env.StopBinaries()
	})
}

func Fuzz_Nosy_Config_StopGraphNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, testName string, rootPath string) {
		env := NewTestConfig(testName, rootPath)
		env.StopGraphNode()
	})
}

// skipping Fuzz_Nosy_Config_applyDefaults__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_Config_deployEigenDAContracts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, testName string, rootPath string) {
		env := NewTestConfig(testName, rootPath)
		env.deployEigenDAContracts()
	})
}

// skipping Fuzz_Nosy_Config_deploySubgraph__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/inabox/deploy.subgraphUpdater

func Fuzz_Nosy_Config_deploySubgraphs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, testName string, rootPath string, startBlock int) {
		env := NewTestConfig(testName, rootPath)
		env.deploySubgraphs(startBlock)
	})
}

func Fuzz_Nosy_Config_genNodeService__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var testName string
		fill_err = tp.Fill(&testName)
		if fill_err != nil {
			return
		}
		var rootPath string
		fill_err = tp.Fill(&rootPath)
		if fill_err != nil {
			return
		}
		var compose testbed
		fill_err = tp.Fill(&compose)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var image string
		fill_err = tp.Fill(&image)
		if fill_err != nil {
			return
		}
		var envFile string
		fill_err = tp.Fill(&envFile)
		if fill_err != nil {
			return
		}
		var ports []string
		fill_err = tp.Fill(&ports)
		if fill_err != nil {
			return
		}

		env := NewTestConfig(testName, rootPath)
		env.genNodeService(compose, name, image, envFile, ports)
	})
}

func Fuzz_Nosy_Config_genService__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var testName string
		fill_err = tp.Fill(&testName)
		if fill_err != nil {
			return
		}
		var rootPath string
		fill_err = tp.Fill(&rootPath)
		if fill_err != nil {
			return
		}
		var compose testbed
		fill_err = tp.Fill(&compose)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var image string
		fill_err = tp.Fill(&image)
		if fill_err != nil {
			return
		}
		var envFile string
		fill_err = tp.Fill(&envFile)
		if fill_err != nil {
			return
		}
		var ports []string
		fill_err = tp.Fill(&ports)
		if fill_err != nil {
			return
		}

		env := NewTestConfig(testName, rootPath)
		env.genService(compose, name, image, envFile, ports)
	})
}

func Fuzz_Nosy_Config_generateBatcherVars__(f *testing.F) {
	f.Fuzz(func(t *testing.T, testName string, rootPath string, ind int, key string, graphUrl string, logPath string) {
		env := NewTestConfig(testName, rootPath)
		env.generateBatcherVars(ind, key, graphUrl, logPath)
	})
}

func Fuzz_Nosy_Config_generateChurnerVars__(f *testing.F) {
	f.Fuzz(func(t *testing.T, testName string, rootPath string, ind int, graphUrl string, logPath string, grpcPort string) {
		env := NewTestConfig(testName, rootPath)
		env.generateChurnerVars(ind, graphUrl, logPath, grpcPort)
	})
}

func Fuzz_Nosy_Config_generateDisperserVars__(f *testing.F) {
	f.Fuzz(func(t *testing.T, testName string, rootPath string, ind int, key string, address string, logPath string, dbPath string, grpcPort string) {
		env := NewTestConfig(testName, rootPath)
		env.generateDisperserVars(ind, key, address, logPath, dbPath, grpcPort)
	})
}

func Fuzz_Nosy_Config_generateEigenDADeployConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, testName string, rootPath string) {
		env := NewTestConfig(testName, rootPath)
		env.generateEigenDADeployConfig()
	})
}

func Fuzz_Nosy_Config_generateEncoderVars__(f *testing.F) {
	f.Fuzz(func(t *testing.T, testName string, rootPath string, ind int, grpcPort string) {
		env := NewTestConfig(testName, rootPath)
		env.generateEncoderVars(ind, grpcPort)
	})
}

func Fuzz_Nosy_Config_generateOperatorVars__(f *testing.F) {
	f.Fuzz(func(t *testing.T, testName string, rootPath string, ind int, name string, key string, churnerUrl string, logPath string, dbPath string, dispersalPort string, retrievalPort string, metricsPort string, nodeApiPort string) {
		env := NewTestConfig(testName, rootPath)
		env.generateOperatorVars(ind, name, key, churnerUrl, logPath, dbPath, dispersalPort, retrievalPort, metricsPort, nodeApiPort)
	})
}

func Fuzz_Nosy_Config_generateRetrieverVars__(f *testing.F) {
	f.Fuzz(func(t *testing.T, testName string, rootPath string, ind int, key string, graphUrl string, logPath string, grpcPort string) {
		env := NewTestConfig(testName, rootPath)
		env.generateRetrieverVars(ind, key, graphUrl, logPath, grpcPort)
	})
}

func Fuzz_Nosy_Config_getKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, testName string, rootPath string, name string) {
		env := NewTestConfig(testName, rootPath)
		env.getKey(name)
	})
}

func Fuzz_Nosy_Config_getKeyString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, testName string, rootPath string, name string) {
		env := NewTestConfig(testName, rootPath)
		env.getKeyString(name)
	})
}

func Fuzz_Nosy_Config_getPaths__(f *testing.F) {
	f.Fuzz(func(t *testing.T, testName string, rootPath string, name string) {
		env := NewTestConfig(testName, rootPath)
		env.getPaths(name)
	})
}

func Fuzz_Nosy_Config_loadPrivateKeys__(f *testing.F) {
	f.Fuzz(func(t *testing.T, testName string, rootPath string) {
		env := NewTestConfig(testName, rootPath)
		env.loadPrivateKeys()
	})
}

// skipping Fuzz_Nosy_Config_updateSubgraph__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/inabox/deploy.subgraphUpdater

func Fuzz_Nosy_EigenDADeployConfig_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *EigenDADeployConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		cfg.MarshalJSON()
	})
}

func Fuzz_Nosy_BatcherVars_getEnvMap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var vars BatcherVars
		fill_err = tp.Fill(&vars)
		if fill_err != nil {
			return
		}

		vars.getEnvMap()
	})
}

func Fuzz_Nosy_ChurnerVars_getEnvMap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var vars ChurnerVars
		fill_err = tp.Fill(&vars)
		if fill_err != nil {
			return
		}

		vars.getEnvMap()
	})
}

func Fuzz_Nosy_Config_IsEigenDADeployed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, testName string, rootPath string) {
		c := NewTestConfig(testName, rootPath)
		c.IsEigenDADeployed()
	})
}

func Fuzz_Nosy_DisperserVars_getEnvMap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var vars DisperserVars
		fill_err = tp.Fill(&vars)
		if fill_err != nil {
			return
		}

		vars.getEnvMap()
	})
}

func Fuzz_Nosy_EncoderVars_getEnvMap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var vars EncoderVars
		fill_err = tp.Fill(&vars)
		if fill_err != nil {
			return
		}

		vars.getEnvMap()
	})
}

func Fuzz_Nosy_Environment_IsLocal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e Environment
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}

		e.IsLocal()
	})
}

func Fuzz_Nosy_OperatorVars_getEnvMap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var vars OperatorVars
		fill_err = tp.Fill(&vars)
		if fill_err != nil {
			return
		}

		vars.getEnvMap()
	})
}

func Fuzz_Nosy_RetrieverVars_getEnvMap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var vars RetrieverVars
		fill_err = tp.Fill(&vars)
		if fill_err != nil {
			return
		}

		vars.getEnvMap()
	})
}

func Fuzz_Nosy_eigenDAOperatorStateSubgraphUpdater_UpdateNetworks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var u eigenDAOperatorStateSubgraphUpdater
		fill_err = tp.Fill(&u)
		if fill_err != nil {
			return
		}
		var n Networks
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var startBlock int
		fill_err = tp.Fill(&startBlock)
		if fill_err != nil {
			return
		}

		u.UpdateNetworks(n, startBlock)
	})
}

func Fuzz_Nosy_eigenDAOperatorStateSubgraphUpdater_UpdateSubgraph__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var u eigenDAOperatorStateSubgraphUpdater
		fill_err = tp.Fill(&u)
		if fill_err != nil {
			return
		}
		var s *Subgraph
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var startBlock int
		fill_err = tp.Fill(&startBlock)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		u.UpdateSubgraph(s, startBlock)
	})
}

func Fuzz_Nosy_eigenDAUIMonitoringUpdater_UpdateNetworks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var u eigenDAUIMonitoringUpdater
		fill_err = tp.Fill(&u)
		if fill_err != nil {
			return
		}
		var n Networks
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var startBlock int
		fill_err = tp.Fill(&startBlock)
		if fill_err != nil {
			return
		}

		u.UpdateNetworks(n, startBlock)
	})
}

func Fuzz_Nosy_eigenDAUIMonitoringUpdater_UpdateSubgraph__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var u eigenDAUIMonitoringUpdater
		fill_err = tp.Fill(&u)
		if fill_err != nil {
			return
		}
		var s *Subgraph
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var startBlock int
		fill_err = tp.Fill(&startBlock)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		u.UpdateSubgraph(s, startBlock)
	})
}

// skipping Fuzz_Nosy_subgraphUpdater_UpdateNetworks__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/inabox/deploy.subgraphUpdater

// skipping Fuzz_Nosy_subgraphUpdater_UpdateSubgraph__ because parameters include func, chan, or unsupported interface: github.com/Layr-Labs/eigenda/inabox/deploy.subgraphUpdater

func Fuzz_Nosy_CallContract__(f *testing.F) {
	f.Fuzz(func(t *testing.T, destination string, signature string, rpcUrl string) {
		CallContract(destination, signature, rpcUrl)
	})
}

func Fuzz_Nosy_CreateNewTestDirectory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, templateName string, rootPath string) {
		CreateNewTestDirectory(templateName, rootPath)
	})
}

func Fuzz_Nosy_GetAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, privateKey string) {
		GetAddress(privateKey)
	})
}

func Fuzz_Nosy_GetLatestBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, rpcUrl string) {
		GetLatestBlockNumber(rpcUrl)
	})
}

func Fuzz_Nosy_GetLatestTestDirectory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, rootPath string) {
		GetLatestTestDirectory(rootPath)
	})
}

func Fuzz_Nosy_PurgeDockertestResources__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *dockertest.Pool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var resource *dockertest.Resource
		fill_err = tp.Fill(&resource)
		if fill_err != nil {
			return
		}
		if pool == nil || resource == nil {
			return
		}

		PurgeDockertestResources(pool, resource)
	})
}

func Fuzz_Nosy_ReadEnv__(f *testing.F) {
	f.Fuzz(func(t *testing.T, filename string) {
		ReadEnv(filename)
	})
}

func Fuzz_Nosy_StartDockertestWithLocalstackContainer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, localStackPort string) {
		StartDockertestWithLocalstackContainer(localStackPort)
	})
}

func Fuzz_Nosy_changeDirectory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string) {
		changeDirectory(path)
	})
}

func Fuzz_Nosy_createDirectory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, name string) {
		createDirectory(name)
	})
}

func Fuzz_Nosy_execBashCmd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, command string) {
		execBashCmd(command)
	})
}

func Fuzz_Nosy_execForgeScript__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var script string
		fill_err = tp.Fill(&script)
		if fill_err != nil {
			return
		}
		var privateKey string
		fill_err = tp.Fill(&privateKey)
		if fill_err != nil {
			return
		}
		var deployer *ContractDeployer
		fill_err = tp.Fill(&deployer)
		if fill_err != nil {
			return
		}
		var extraArgs []string
		fill_err = tp.Fill(&extraArgs)
		if fill_err != nil {
			return
		}
		if deployer == nil {
			return
		}

		execForgeScript(script, privateKey, deployer, extraArgs)
	})
}

func Fuzz_Nosy_execYarnCmd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var command string
		fill_err = tp.Fill(&command)
		if fill_err != nil {
			return
		}
		var args []string
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}

		execYarnCmd(command, args...)
	})
}

func Fuzz_Nosy_genTelemetryServices__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var compose testbed
		fill_err = tp.Fill(&compose)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var image string
		fill_err = tp.Fill(&image)
		if fill_err != nil {
			return
		}
		var volumes []string
		fill_err = tp.Fill(&volumes)
		if fill_err != nil {
			return
		}

		genTelemetryServices(compose, name, image, volumes)
	})
}

func Fuzz_Nosy_readFile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, name string) {
		readFile(name)
	})
}

func Fuzz_Nosy_writeEnv__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var envMap map[string]string
		fill_err = tp.Fill(&envMap)
		if fill_err != nil {
			return
		}
		var filename string
		fill_err = tp.Fill(&filename)
		if fill_err != nil {
			return
		}

		writeEnv(envMap, filename)
	})
}

func Fuzz_Nosy_writeFile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, name string, d2 []byte) {
		writeFile(name, d2)
	})
}
