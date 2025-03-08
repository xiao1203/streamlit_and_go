FROM python:3.9-slim

WORKDIR /app

# 必要なPythonパッケージのインストール
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# アプリケーションのソースコードをコピー
COPY . .

# PYTHONPATHに proto_files ディレクトリを追加
ENV PYTHONPATH="/app/proto_files:${PYTHONPATH}"

# ポート8501を公開（Streamlitのデフォルトポート）
EXPOSE 8501

# Streamlitアプリを起動
CMD ["streamlit", "run", "streamlit_client.py", "--server.address=0.0.0.0"]
