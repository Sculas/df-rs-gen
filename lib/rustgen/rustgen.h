#include <stdint.h>

typedef struct Buffer {
  int32_t len;
  uint8_t *data;
} Buffer;

Buffer gen_chunk(int seed, int chunk_x, int chunk_z);
void free_chunk_data(Buffer buf);
