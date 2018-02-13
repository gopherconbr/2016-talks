while(true) {

	// Read doesn't wait for some input
	ssize_t n = read(fd, buffer, sizeof(buffer));

	if (n > 0) {
	    doSomething(buffer, n);
	}
	
    if (n <= 0) {
		if (errno == EWOULDBLOCK) {
			sleep(1);
			continue;
		}
		break;
	}
}
