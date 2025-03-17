use actix_web::{get, HttpResponse, Responder};

#[get("/v1/deploy")]
pub async fn deploy() -> impl Responder {
    println!("Service::Deploy");
    HttpResponse::Ok().body("SolanaDeployer :: Service :: Deploy")
}

// use solana_client::rpc_client::RpcClient;
// use solana_sdk::{
//     pubkey::Pubkey,
//     signature::{Keypair, Signer},
//     system_program,
//     transaction::Transaction,
// };

// fn main() {
//     // Connect to Solana devnet
//     let rpc_url = "https://api.devnet.solana.com";
//     let client = RpcClient::new(rpc_url.to_string());
//
//     // Generate a new keypair
//     let new_account = Keypair::new();
//
//     // Request an airdrop (1 SOL)
//     let airdrop_amount = 1_000_000_000; // 1 SOL in lamports
//     match client.request_airdrop(&new_account.pubkey(), airdrop_amount) {
//         Ok(signature) => {
//             println!("Airdrop requested::: {}", signature);
//         }
//         Err(e) => {
//             eprintln!("Error requesting airdrop:::: {:?}", e);
//             return;
//         }
//     }
//
//     // Wait for airdrop confirmation
//     let recent_blockhash = client.get_recent_blockhash().expect("Failed to get recent blockhash");
//     let signers = [&new_account];
//
//     let tx = Transaction::new_signed_with_payer(
//         &[],
//         Some(&new_account.pubkey()),
//         &signers,
//         recent_blockhash.0,
//     );
//
//     match client.send_and_confirm_transaction(&tx) {
//         Ok(confirmation) => {
//             println!("Transaction confirmed: {:?}", confirmation);
//         }
//         Err(e) => {
//             eprintln!("Error confirming transaction: {:?}", e);
//             return;
//         }
//     }
//
//     // Print the public key of the new account
//     println!("New account public key:: {}", new_account.pubkey());
// }
