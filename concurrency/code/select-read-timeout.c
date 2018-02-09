// initialize readfds with fd1 and fd2

// setup to five seconds
tv.tv_sec = 5;
tv.tv_usec = 0;

while(true) {
	// select waits for some input on fd1 or fd2
	retval = select(nfds, readfds, NULL, NULL, &tv);
	if (retval == -1) {
		break;
	}
	
	if (retval == 0) {
        printf("No data within five seconds.\n");
		continue;
	}

    if (FD_ISSET(fd1, &readfds[0]) {
		ssize_t n = read(fd1, buffer, sizeof(buffer));
		if (n > 0) {
    		doSomething(buffer, n);
		}
	}

    if (FD_ISSET(fd2, &readfds[1]) {
		ssize_t n = read(fd2, buffer, sizeof(buffer));
		if (n > 0) {
    		doSomething(buffer, n);
		}
	}
}
