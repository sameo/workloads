#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int main(int argc, char * argv[])
{
	extern char *optarg; extern int optind;
	int loop, opt, i, ret;
	char *loop_string = "10000";
	
	while ((opt = getopt(argc, argv, "l:")) != -1) {
		switch (opt) {
		case 'l':
			loop_string = optarg;
			break;
		default:
			fprintf(stderr, "Usage: %s [-ilw] [file...]\n", argv[0]);
			exit(EXIT_FAILURE);
		}
	}

	loop = atoi(loop_string);
	
	for (i = 0; i < loop; i++) {
		ret = fork();
		if (!ret) {
			exit(0);
		}
	}
}
