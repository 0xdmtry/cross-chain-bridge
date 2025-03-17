mod services;

use actix_web::{App, HttpServer};
use crate::services::deploy;

#[actix_rt::main]
async fn main() -> std::io::Result<()> {
    println!("SolanaDeployer :: Main");

    HttpServer::new(move || {
        App::new()
            .service(deploy)
    })
        .bind(("0.0.0.0", 8000))?
        .run()
        .await
}

