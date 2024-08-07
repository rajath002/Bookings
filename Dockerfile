# Use the official Ubuntu 22.04 image as the base image
FROM ubuntu:22.04

# Set environment variables
ENV DEBIAN_FRONTEND=noninteractive
ENV TZ=Etc/UTC

# Set environment variables for PostgreSQL
ENV POSTGRES_USER=postgresadmin
ENV POSTGRES_PASSWORD=postgresadmin
ENV POSTGRES_DB=bookings

# Install dependencies and set timezone
RUN apt-get update && \
    apt-get install -y wget gnupg2 lsb-release software-properties-common tzdata && \
    ln -fs /usr/share/zoneinfo/$TZ /etc/localtime && \
    dpkg-reconfigure --frontend noninteractive tzdata

# Add PostgreSQL repository
RUN echo "deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list && \
    wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | apt-key add -

# Install PostgreSQL, Golang, and other necessary packages
RUN apt-get update && \
    apt-get install -y postgresql-14 postgresql-client-14 git curl

# Install Golang
RUN curl -O https://storage.googleapis.com/golang/go1.21.0.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz

# Set Go environment variables
ENV PATH=$PATH:/usr/local/go/bin
ENV GOPATH=/go
ENV PATH=$PATH:$GOPATH/bin

# Switch to postgres user and set up the database
USER postgres
RUN /etc/init.d/postgresql start && \
    psql --command "CREATE USER ${POSTGRES_USER} WITH SUPERUSER PASSWORD '${POSTGRES_PASSWORD}';" && \
    createdb -O ${POSTGRES_USER} ${POSTGRES_DB}

# Switch back to root user
USER root

# Create a working directory for the application
WORKDIR /app/bookings

# Copy local project files to the container
COPY . /app/bookings

# Install Soda CLI
RUN curl -sL https://github.com/gobuffalo/pop/releases/download/v5.3.3/pop_5.3.3_linux_amd64.tar.gz | tar xz -C /usr/local/bin

# Install Go dependencies
RUN go mod tidy

RUN chmod 777 /app/bookings/run.sh
# Build the Go application
# RUN go build -o run.sh .

# Run Soda migrations
# RUN /etc/init.d/postgresql start && \
#     soda migrate

# Expose the necessary ports
EXPOSE 5432 8080

# Run the application
# CMD ["/app/run.sh"]
