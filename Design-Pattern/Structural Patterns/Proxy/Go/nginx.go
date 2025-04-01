package main

type Nginx struct {
	application        *Application
	maxAllowedRequests int
	rateLimiter        map[string]int
}

func newNginxServer() *Nginx {
	return &Nginx {
		application: &Application{},
		maxAllowedRequests: 2,
		rateLimiter: make(map[string]int),
	}
}

func (n *Nginx) handleRequest(url, method string) (int, string) {
	allowed := n.checkRateLimiting(url)
	if !allowed {
		return 403, "Not Allowed"
	}

	return n.application.handleRequest(url, method)
}

func (n *Nginx) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}
	if n.rateLimiter[url] > n.maxAllowedRequests {
		return false
	}

	n.rateLimiter[url]++
	return true
}
