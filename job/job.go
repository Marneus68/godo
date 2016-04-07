package job

// Structure describing a godo job
//
// Name : name of the job
// Tags : job tags expected by this job to be run by a slave or servant instance
// Command : command executed by the job
type Job struct {
	Name    string
	Tags    []string
	Command string
}

// Creates a new empty job
func NewJob() *Job {
	j := new(Job)
	j.Name = ""
	j.Tags = make([]string, 0)
	j.Command = ""
	return j
}

// Creates a new job from a file
func NewJobFromFile(path string) *Job {
	j := new(Job)
	j.Name = ""
	j.Tags = make([]string, 0)
	j.Command = ""
	return j
}

// Read job from file
func (job Job) ReadFromFile(path string) {

}

// Save job to file
func (job Job) SaveToFile(path string) {

}
