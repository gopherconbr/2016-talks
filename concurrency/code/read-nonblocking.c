unsigned char buffer[1024];
int fd;

fd = open("books-to-read.txt", O_RDWR | O_NONBLOCK);

while(true) {

	// Read doesn't wait for some input
	ssize_t n = read(fd, buffer, sizeof(buffer));

	if (n > 0) {
	    doSomething(buffer, n);
	}
	
    if (n <= 0) {
		if (errno == EWOULDBLOCK) {
			continue;
		}
		break;
	}
}

close(fd);
