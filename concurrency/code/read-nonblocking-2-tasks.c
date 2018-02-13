while(task1_reading && task2_reading) {

	if (task1_reading) {
		ssize_t n1 = read(fd1, buffer1, sizeof(buffer1));
		if (n > 0) {
	    	doSomething(buffer1, n1);
		}
	
    	if (n1 <= 0 && errno != EWOULDBLOCK) {
			task1_reading = false;
		}
	}

	if (task2_reading) {
		ssize_t n2 = read(fd2, buffer2, sizeof(buffer2));
		if (n > 0) {
	    	doSomething(buffer2, n2);
		}
	
    	if (n2 <= 0 && errno != EWOULDBLOCK) {
			task2_reading = false;
		}
	}
}
