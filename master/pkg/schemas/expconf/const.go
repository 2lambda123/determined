package expconf

// Configuration constants for task name generator.
const (
	TaskNameGeneratorWords = 3
	TaskNameGeneratorSep   = "-"
)

// Default task environment docker image names.
const (
	CPUImage  = "determinedai/environments-dev:py-3.10-pytorch-2.0-tf-2.11-cpu-7e044b9"
	CUDAImage = "determinedai/environments-dev:cuda-11.8-pytorch-2.0-tf-2.11-gpu-7e044b9"
	ROCMImage = "determinedai/environments-dev:rocm-5.0-pytorch-1.10-tf-2.7-rocm-7e044b9"
)
