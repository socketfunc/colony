#[no_mangle]
pub fn add(a: i32, b: i32) -> i32 {
    a + b
}

#[no_mangle]
pub fn hello() -> *const u8 {
    b"Hello World\0".as_ptr()
}
