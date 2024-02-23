FROM rust:1.76.0
WORKDIR /back-end/
COPY /back-end/ /back-end/
RUN cargo install --path .
CMD ["back-end"]