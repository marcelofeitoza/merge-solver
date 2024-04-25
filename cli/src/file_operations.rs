use tokio::fs::File;
use tokio::io::{AsyncReadExt, AsyncWriteExt};
use anyhow::Result;

pub async fn read_file(path: String) -> Result<String> {
    let mut file = File::open(path).await?;
    let mut contents = String::new();
    file.read_to_string(&mut contents).await?;
    Ok(contents)
}

pub async fn write_file(path: &str, content: &str) -> Result<()> {
    let mut file = File::create(path).await?;
    file.write_all(content.as_bytes()).await?;
    Ok(())
}
