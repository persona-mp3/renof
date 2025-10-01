#include <errno.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/stat.h>

void reader(const char fname[]) {
  FILE *fptr;
  fptr = fopen(fname, "r+");
  if (!fptr) {
    printf("error_n -> %d\n", errno);
    perror("error: ");
    return;
  }

  // so there are a couple of unknowns here,
  // we need to know the size of the file, before allocating space for its
  // content it could be 500bytes or more , before we can hash the content
  char line[255];
  while (fgets(line, sizeof(line) + 1, fptr)) {
    printf("%s", line);
  }
  fclose(fptr);

  return;
}

int file_reader(const char fname[]) 
{

  // https://man7.org/linux/man-pages/man2/stat.2.html
  struct stat file_stat;
  if (stat(fname, &file_stat) == -1) {
    printf("error_n -> %d\n", errno);
    perror("error: ");
    return EXIT_FAILURE;
  }

  size_t file_size = file_stat.st_size;
  printf("size -> %lli\n", file_stat.st_size);
  char *buffer = malloc(file_size + 1);
  if (!buffer){
    printf("error in allocating memory for file of size -> %lu\n", file_size);
    printf("e -> %d\n", errno);
    perror("e :");
    return EXIT_FAILURE;
  }

  FILE *fptr;
  fptr = fopen(fname, "rb");
  if (!fptr) {
    printf("error in opening file %s\n", fname);
    printf("e -> %d\n", errno);
    perror("e :");
    free(buffer);
    return EXIT_FAILURE;
  }

  printf("reading from file\n");
  // fread(*buffer, element_size, n_elements,src)
  size_t bytes_read = fread(buffer, 1, file_size,fptr);
  fclose(fptr);

  buffer[bytes_read] = '\0';

  free(buffer);
  return EXIT_SUCCESS;
}

int main() {
  printf("renoc\n\n");
  char fname[30];
  printf("file_name -> ");
  scanf("%s", fname);

  file_reader(fname);
  return 0;
}
