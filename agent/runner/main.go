package runner

type Runner interface {
	ImagePull()

	RunContainer()

	GetContainerLogs()
}
