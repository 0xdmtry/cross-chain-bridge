# Use the official Rust image as a base
FROM rust:1.71.1

# Install cargo-watch
RUN cargo install cargo-watch

# Set the working directory inside the container
WORKDIR /usr/src/app

# Copy the Cargo.toml and Cargo.lock files into the container
COPY bridge-solana-deployer/Cargo.toml bridge-solana-deployer/Cargo.lock bridge-solana-deployer/.env ./


# This step caches your dependencies
RUN mkdir src/ && \
    echo "fn main() {}" > src/main.rs && \
    cargo build --release && \
    rm -f target/release/deps/bridge-solana-deployer*

# Copy the source code into the container
COPY bridge-solana-deployer/src ./src

EXPOSE 8000

# Command to run when starting the container
CMD ["cargo", "watch", "-x", "run"]
