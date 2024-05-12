# FROM golang:1.22.2-bullseye
# ARG USERNAME=mochi
# ARG GROUPNAME=mochi
# ARG UID=1000
# ARG GID=1000
# RUN groupadd -g ${GID} ${GROUPNAME} && \
#     useradd -s /bin/bash -u ${UID} -g ${GID} ${USERNAME} -d /home/${USERNAME} -m -k /etc/skel && \
#     mkdir -p /home/${USERNAME}/
#     # chown -R ${USERNAME}:${GROUPNAME} /home/${USERNAME} && \
#     # chmod -R 750 /home/${USERNAME}
#     # mkdir -p /src && \
#     # chown -R ${USERNAME}:${GROUPNAME} /src
# USER ${USERNAME}
# WORKDIR /home/${USERNAME}/
# CMD [ "/bin/bash" ]

FROM golang:1.22.2-bullseye

ARG USERNAME=mochi
ARG GROUPNAME=mochi
ARG UID=1000
ARG GID=1000

# 非ルートユーザーを作成する
RUN groupadd -g ${GID} ${GROUPNAME} && \
    useradd -u ${UID} -g ${GROUPNAME} -m ${USERNAME} -s /bin/bash

# ユーザーを切り替える
USER ${USERNAME}

# ワーキングディレクトリを設定
WORKDIR /home/${USERNAME}/go/src/app

# 必要なライブラリやツールのインストールなど、追加のセットアップがあればここに記述する

# ターミナルを使う場合には、bashを起動する
CMD ["/bin/bash"]
