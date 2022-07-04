require("dotenv").config();
const {Client, Intents, MessageEmbed} = require('discord.js');

const client = new Client({
    intents: [
        Intents.FLAGS.GUILDS,
        Intents.FLAGS.GUILD_MESSAGES
    ]
});

client.on('ready', () => {
    console.log(`Logged in as ${client.user.tag}!`);
});

client.on('messageCreate', message => {
    if (message.content === 'b') {
        const sentMessage = sendStartMessage(message)
    }
});

// メッセージを送信します
const sendStartMessage = async (message) => {
    const exampleEmbed = {
        color: 0x0099ff,
        title: '⚔️ Battle Start ⚔️',
        description: '\nねだるな、勝ち取れ\n',
        // サムネイルは後で追加する
        // thumbnail: {
        //     url: 'https://i.imgur.com/AfFp7pu.png',
        // },
        fields: [
            {
                name: '主催者',
                value: `${message.author}`,
            },
            {
                name: '勝者',
                value: '1名',
                inline: false,
            },
            {
                name: 'エントリー',
                value: '⚔️のリアクション',
                inline: false,
            },
            {
                name: '試合開始',
                value: 'メッセージ送信から2分後',
                inline: false,
            },
        ],
        timestamp: new Date(),
    };

    const sentMessage = await message.channel.send({embeds: [exampleEmbed]})

    try {
        await sentMessage.react("⚔️");
    } catch (error) {
        console.error("emoji failed to react")
    }

    return sentMessage;
}

client.login(process.env.APP_BOT_TOKEN);