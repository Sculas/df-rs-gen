extern crate voxel_worldgen;
extern crate rand;
extern crate nalgebra;
extern crate noise;

use std::os::raw::c_int;

use voxel_worldgen::generators;
use voxel_worldgen::rnd::RngBuilder;

use nalgebra::Pnt2;

#[repr(C)]
pub struct Buffer {
    len: i32,
    data: *mut u8,
}

#[no_mangle]
pub extern "C" fn gen_chunk(seed: c_int, chunk_x: c_int, chunk_z: c_int) -> Buffer  {
    let mut seed_rng = RngBuilder::init().mix(seed as u64).build();
    let world_gen_state = generators::vanilla::WorldGeneratorState::new(&mut seed_rng);

    let chunk_pos = Pnt2::new(chunk_x, chunk_z);
    let mut chunk = generators::vanilla::generate_chunk(&world_gen_state, chunk_pos).data;

    let mut buf = chunk.into_boxed_slice();
    let data = buf.as_mut_ptr();
    let len = buf.len() as i32;
    std::mem::forget(buf);
    Buffer { len, data }
}

#[no_mangle]
pub extern "C" fn free_chunk_data(buf: Buffer) {
    let s = unsafe { std::slice::from_raw_parts_mut(buf.data, buf.len as usize) };
    let s = s.as_mut_ptr();
    unsafe {
        Box::from_raw(s);
    }
}

// TODO: In Go, to convert:
// const BLOCK_AIR: u8 = 0;
// const BLOCK_STONE: u8 = 1;
// const BLOCK_GRASS: u8 = 2;
// const BLOCK_DIRT: u8 = 3;
// const BLOCK_FLOWING_WATER: u8 = 8;
// const BLOCK_WATER: u8 = 9;
// const BLOCK_FLOWING_LAVA: u8 = 10;
// const BLOCK_LAVA: u8 = 11;