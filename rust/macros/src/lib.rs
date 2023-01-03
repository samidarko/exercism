#[macro_export(local_inner_macros)]
macro_rules! hashmap {
    (@single $($x:tt)*) => (());
    (@count $($rest:expr),*) => (<[()]>::len(&[$(hashmap!(@single $rest)),*]));

    ($($key:expr => $value:expr,)+) => { hashmap!($($key => $value),+) };
    ($($key:expr => $value:expr),*) => {
        {
            let _cap = hashmap!(@count $($key),*);
            let mut _hm = ::std::collections::HashMap::with_capacity(_cap);
            $(
                _hm.insert($key, $value);
            )*
            _hm
        }
    };
}
