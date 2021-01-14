FROM ruby:2.6

ENV LANG=ja_JP.UTF-8
ENV TZ=Asia/Tokyo
ENV ROOT=/myapp \
    GEM_HOME=/bundle \
    BUNDLE_PATH=$GEM_HOME
ENV BUNDLE_BIN=$BUNDLE_PATH/bin
ENV PATH /app/bin:$BUNDLE_BIN:$PATH


WORKDIR $ROOT

RUN curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add - \
    && echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list \
    && apt-get update -qq \
    && apt-get install -y nodejs yarn

COPY Gemfile $ROOT
COPY Gemfile.lock $ROOT

RUN bundle install -j4
#　不要ファイル削除
RUN rm -rf /usr/local/bundle/cache/* /usr/local/share/.cache/* /var/cache/* /tmp/*

COPY . $ROOT

# Add a script to be executed every time the container starts.
COPY entrypoint.sh /usr/bin/
RUN chmod +x /usr/bin/entrypoint.sh
ENTRYPOINT ["sh", "/usr/bin/entrypoint.sh"]
EXPOSE 3000