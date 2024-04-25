mod file_operations;
mod http;

use clap::{Parser, Subcommand};
use dialoguer::Confirm;
use dotenv::dotenv;
use file_operations::{read_file, write_file};
use http::{post_data, prettify};

#[derive(Subcommand, Debug)]
enum Commands {
    Merge { old: String, new: String },
}

#[derive(Parser, Debug)]
#[clap(
    version = "1.0",
    author = "DeployBuddy",
    about = "DeployBuddy's Merge Solver is a tool that makes it easier to automatically merge code files without having to manually edit them."
)]
struct Args {
    #[command(subcommand)]
    cmd: Commands,
}

#[tokio::main]
async fn main() -> anyhow::Result<()> {
    dotenv().ok();
    let args = Args::parse();

    match args.cmd {
        Commands::Merge { old, new } => {
            let old_content = read_file(old).await?;
            let new_content = read_file(new).await?;

            let mut payload = http::Payload {
                old: old_content,
                new: new_content,
                rejected: None,
            };

            let mut attempt = 1;
            loop {
                println!("Merge attempt {}: sending data to server...", attempt);
                let response = post_data(payload.clone()).await?;
                let pretty_response = prettify(&response.content);
                println!("Server responded with proposed merge:\n{}", pretty_response);
                write_file("output.txt", &pretty_response).await?;

                if Confirm::new().with_prompt("Do you accept this merge?").interact()? {
                    println!("Merge accepted. The merged content has been saved to 'output.txt'.");
                    break;
                } else if Confirm::new().with_prompt("Do you want to rejected a new merge proposal?").interact()? {
                    attempt += 1;
                    payload.rejected = Some(response.content);
                } else {
                    println!("Merge rejected by user. No further action will be taken.");
                    break;
                }
            }
        }
    }

    Ok(())
}
