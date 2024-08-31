FROM golang:1.22.2-bullseye

ARG USERNAME=mochi
ARG GROUPNAME=mochi
ARG UID=1000
ARG GID=1000

# 非ルートユーザーを作成する
RUN groupadd -g ${GID} ${GROUPNAME} && \
    useradd -u ${UID} -g ${GROUPNAME} -m ${USERNAME} -s /bin/bash && \
    mkdir -p /home/${USERNAME}/go/src/app && \
    chown -R ${USERNAME}:${GROUPNAME} /home/${USERNAME} && \
    chmod -R 755 /home/${USERNAME}/go/src/app

# ユーザーを切り替える
USER ${USERNAME}

# ワーキングディレクトリを設定
WORKDIR /home/${USERNAME}/go/src/app

# 必要なライブラリやツールのインストールなど、追加のセットアップがあればここに記述する

# ターミナルを使う場合には、bashを起動する
CMD ["/bin/bash"]
