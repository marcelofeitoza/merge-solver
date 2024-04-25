use serde::{Deserialize, Serialize};
use anyhow::Result;

#[derive(Serialize, Deserialize, Debug)]
pub struct Response {
    pub content: String,
    pub role: String,
}

#[derive(Serialize, Debug, Clone)]
pub struct Payload {
    pub old: String,
    pub new: String,
    pub rejected: Option<String>,
}

pub async fn post_data(payload: Payload) -> Result<Response> {
    let base_url = dotenv::var("BASE_URL")?;
    let client = reqwest::Client::new();
    let response = client
        .post(format!("{}/merge", base_url))
        .json(&payload)
        .send()
        .await?
        .json()
        .await?;

    println!("{:#?}", response);

    Ok(response)
}

pub fn prettify(content: &str) -> String {
    content.trim().lines()
        .filter(|line| !line.is_empty())
        .collect::<Vec<&str>>()
        .join("\n")
}
