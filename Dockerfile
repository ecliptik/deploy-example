FROM ruby:2.3

LABEL maintainer="ecliptik@gmail.com"

WORKDIR /app

COPY . /app
RUN bundle install

EXPOSE 9292

ENTRYPOINT ["bundle"]
CMD ["exec", "rackup", "-o", "0.0.0.0"]
