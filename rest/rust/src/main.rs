use warp::Filter;
use serde::{Deserialize, Serialize};

#[derive(Deserialize)]
struct Numbers {
    values: Vec<f64>,
}

#[derive(Serialize)]
struct ComputeResult {
    sum: f64,
    average: f64,
    product: f64,
}

async fn compute_handler(numbers: Numbers) -> Result<warp::reply::Json, warp::Rejection> {
    let sum: f64 = numbers.values.iter().sum();
    let product: f64 = numbers.values.iter().product();
    let average: f64 = sum / numbers.values.len() as f64;

    Ok(warp::reply::json(&ComputeResult { sum, average, product }))
}

#[tokio::main]
async fn main() {
    let compute = warp::post()
        .and(warp::path("compute"))
        .and(warp::body::json())
        .and_then(compute_handler);

    warp::serve(compute).run(([127, 0, 0, 1], 8080)).await;
}
