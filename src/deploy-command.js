require("dotenv").config();
const {SlashCommandBuilder} = require('@discordjs/builders'); //SlashCommandBuilderを読み込む
const {REST} = require('@discordjs/rest'); //RESTを読み込む
const {Routes} = require('discord-api-types/v9'); //Routesを読み込む
const client_id = process.env.CLIENT_ID //クライアントIDをclient_idに代入
const guild_id = '965048644480950282' //サーバーIDをguild_idに代入
const token = process.env.APP_BOT_TOKEN //トークン

const commands = [
    new SlashCommandBuilder().setName('wallet').setDescription('ウォレットアドレスの入力'), //pingコマンド
]
    .map(command => command.toJSON());

const rest = new REST({version: '9'}).setToken(token);

rest.put(Routes.applicationGuildCommands(client_id, guild_id), {body: commands})
    .then(() => console.log('Successfully registered application commands.'))
    .catch(console.error); //指定したサーバーにコマンドを登録・更新