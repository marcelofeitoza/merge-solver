# Use the official Rust image to create a build artifact.
FROM rust:latest as builder

WORKDIR /app

# Copy the Rust project files first to leverage Docker cache.
COPY ./Cargo.toml ./Cargo.lock ./

# Create a dummy source file to build the dependencies layer.
RUN mkdir src \
    && echo "fn main() {}" > src/main.rs \
    && cargo build --release \
    && rm -f target/release/deps/cli*

# Copy the rest of the application source code.
COPY . .

# Build the binary.
RUN cargo build --release

# Use the debian:buster-slim image for a small, final production image.
FROM debian:buster-slim
RUN apt-get update \
    && apt-get dist-upgrade -y \
    && apt-get install -y ca-certificates \
    && update-ca-certificates

WORKDIR /root/

# Copy the pre-built binary file from the previous stage.
COPY --from=builder /app/target/release/cli .

# Command to run the executable.
CMD ["./cli"]
