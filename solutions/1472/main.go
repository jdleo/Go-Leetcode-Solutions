package main

type BrowserHistory struct {
	homepage string
	front    []string
	back     []string
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{homepage, []string{}, []string{}}
}

func (this *BrowserHistory) Visit(url string) {
	// clear backwards history
	this.back = []string{}
	// push to front
	this.front = append(this.front, url)
}

func (this *BrowserHistory) Back(steps int) string {

}

func (this *BrowserHistory) Forward(steps int) string {

}
