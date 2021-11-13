package runner

import "sync"

var wg sync.WaitGroup

//Running both api and cli
func RunInDefaultMode() {
	wg = sync.WaitGroup{}
	wg.Add(2)

	go func() {
		err := RunCli()
		if err != nil {
			wg.Done()
			return
		}
	}()

	go func() {
		err := RunAPI()
		if err != nil {
			wg.Done()
			return
		}
	}()

	wg.Wait()
}

//Running only cli
func RunCli() error {
	return nil
}

//Running only api
func RunAPI() error {
	return nil
}
