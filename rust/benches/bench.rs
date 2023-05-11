#![feature(test)]
extern crate rmp_serde;
extern crate rust_serde_benchmarks;
extern crate serde;
extern crate serde_json;
extern crate test;

use rust_serde_benchmarks::IngestData;
use serde::{Deserialize, Serialize};
use test::Bencher;

#[bench]
fn clone(b: &mut Bencher) {
    let log = IngestData::new();

    let size = ::std::mem::size_of_val(&log);

    b.bytes = size as u64;

    b.iter(|| log.clone());
}

#[bench]
fn serde_json_serialize(b: &mut Bencher) {
    let log = IngestData::new();
    let mut buf = Vec::new();

    serde_json::to_writer(&mut buf, &log).unwrap();
    b.bytes = buf.len() as u64;
    // let size = ::std::mem::size_of_val(&log);
    // b.bytes = size as u64;

    b.iter(|| {
        buf.clear();
        serde_json::to_writer(&mut buf, &log).unwrap();
    });
}

#[bench]
fn serde_json_deserialize(b: &mut Bencher) {
    let log = IngestData::new();
    let json = serde_json::to_string(&log).unwrap();

    b.bytes = json.len() as u64;
    // let size = ::std::mem::size_of_val(&log);
    // b.bytes = size as u64;

    b.iter(|| serde_json::from_str::<IngestData>(&json).unwrap());
}

#[bench]
fn rmp_serde_serialize(b: &mut Bencher) {
    let mut buf = Vec::new();
    let log = IngestData::new();
    log.serialize(&mut ::rmp_serde::Serializer::new(&mut buf))
        .unwrap();
    b.bytes = buf.len() as u64;
    // let size = ::std::mem::size_of_val(&log);
    // b.bytes = size as u64;

    b.iter(|| {
        buf.clear();
        log.serialize(&mut ::rmp_serde::Serializer::new(&mut buf))
            .unwrap();
    });
}

#[bench]
fn rmp_serde_deserialize(b: &mut Bencher) {
    let mut buf = Vec::new();
    let log = IngestData::new();

    log.serialize(&mut ::rmp_serde::Serializer::new(&mut buf))
        .unwrap();
    b.bytes = buf.len() as u64;
    // let size = ::std::mem::size_of_val(&log);
    // b.bytes = size as u64;
    // let slice = buf.as_slice();
    b.iter(|| {
        let mut decoder = ::rmp_serde::Deserializer::new(&*buf);
        let _log: IngestData = Deserialize::deserialize(&mut decoder).unwrap();
        // -        let _log: Log = Deserialize::deserialize(&mut decoder).unwrap();
        // let _outlog: IngestData = rmp_serde::from_slice(slice).unwrap();
    });
}
