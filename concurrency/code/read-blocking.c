#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>

unsigned char buffer[1024];
int fd;

fd = open("books-to-read.txt", O_RDWR);

while(true) {

	// Read waits for some input
	ssize_t n = read(fd, buffer, sizeof(buffer));

    if (n <= 0) break;

	doSomething(buffer, n);
}

close(fd);
