package expconf

// Configuration constants for task name generator.
const (
	TaskNameGeneratorWords = 3
	TaskNameGeneratorSep   = "-"
)

// Default task environment docker image names.
const (
	CPUImage  = "determinedai/environments-dev:py-3.8-pytorch-1.10-lightning-1.5-tf-2.8-cpu-136a60e"
	CUDAImage = "determinedai/environments-dev:cuda-11.3-pytorch-1.10-lightning-1.5-tf-2.8-gpu-136a60e"
	ROCMImage = "determinedai/environments-dev:rocm-4.2-pytorch-1.9-tf-2.5-rocm-136a60e"
)
