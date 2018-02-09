type struct {
	int fd;
	void* userdata;
	void (*callback)(unsigned char* buffer, int n, void* userdata);
} Context;

void func doSomething1(unsigned char* buffer, int n, void* userdata) {...}
void func doSomething2(unsigned char* buffer, int n, void* userdata) {...}

context[0]->fd = fd1; context[0]->callback = doSomething1;
context[1]->fd = fd2; context[1]->callback = doSomething2;

void handleConnections() {
	while(true) {
		retval = select(nfds, readfds, NULL, NULL, NULL);
		if (retval == -1) break;
		for (int i=0; i<nfds; i++) {
    		if (FD_ISSET(contexts[i]>fd, &readfds[i]) {
				ssize_t n = read(context[i]->fd, buffer, sizeof(buffer));
				if (n > 0) context[i]->callback(buffer, n, context[i]->userdata);
			}
		}
	}
}
