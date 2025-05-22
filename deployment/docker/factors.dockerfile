# Etapa 1: build
FROM node:20-alpine AS builder

# Diretório de trabalho
WORKDIR /app

# Copia package.json e instala dependências
COPY ./services/factors/package*.json ./
RUN npm install

# Copia o restante do código
COPY . .

# Transpila os arquivos TypeScript para JavaScript
RUN npm run build

# Etapa 2: imagem final de produção
FROM node:20-alpine

# Diretório da aplicação
WORKDIR /app

# Copia apenas o necessário do builder
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/dist ./dist
COPY ./pkg/services/factors ./pkg/services/factors

# Porta do serviço gRPC
EXPOSE 50052

# Comando de execução
CMD ["node", "dist/server.js"]