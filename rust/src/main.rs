extern crate rmp_serde;
extern crate rust_serde_benchmarks;
extern crate serde;
extern crate serde_json;

use std::borrow::Cow;
use std::time::Instant;

use rust_serde_benchmarks::{IngestData, LogMsg};
use serde::Deserialize;

fn main() {
    // let mut buf = Vec::new();
    let log = LogMsg {
        svc_id: Cow::from("deadbeef"),
        endpoint: Cow::from("sometestendpoint"),
        msg: Vec::from("a").repeat(131072),
    };
    // let m = [97; 131072];
    // let m = Vec::from("a").repeat(131072).as_slice();
    let ingest = IngestData {
        typ: Cow::from("log"),
        source: Cow::from("xqd"),
        timestamp: Cow::from("4:00:00"),
        msg: log,
        // msg: Cow::from("a".to_string().repeat(131072)),
        // msg: Cow::from(&m[..]),
    };

    // log.serialize(&mut ::rmp_serde::Serializer::new(&mut buf))
    //     .unwrap();
    let buf = rmp_serde::encode::to_vec_named(&ingest).unwrap();

    // let size = ::std::mem::size_of_val(&log);
    // b.bytes = size as u64;
    // let decoder = ::rmp_serde::Deserializer::new(&*buf);
    println!("buf size: {}", buf.len());
    // println!("buf: {:?}", buf);
    // let slice = buf.as_slice();
    let time_start = Instant::now();
    for _i in 0..1000000 {
        let mut decoder = ::rmp_serde::Deserializer::new(&*buf);
        let _log: IngestData = Deserialize::deserialize(&mut decoder).unwrap();
        // let _outlog: IngestData = rmp_serde::from_slice(&slice).unwrap();
    }
    let elapsed_time = time_start.elapsed();
    println!("Duration: {}", elapsed_time.as_nanos());
}
