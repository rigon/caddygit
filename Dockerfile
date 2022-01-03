# Build caddy with git module
FROM caddy:builder-alpine AS builder
RUN xcaddy build \
    --with github.com/rigon/caddygit/module/git@docker
