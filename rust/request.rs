use reqwest::{Client, Error};

async fn make_request() -> Result<(), Error> {
    let client = Client::new();

    let res = client
        .get("https://api.github.com/users/octocat")
        .send()
        .await?;

    println!("Status: {}", res.status());
    println!("Headers:\n{}", res.headers());
    println!("Body:\n{}", res.text().await?);

    Ok(())
}

#[tokio::main]
async fn main() -> Result<(), Error> {
    make_request().await
}
