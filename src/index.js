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

client.on('messageCreate', async message => {
    if (message.content === 'b') {
        const entryMessage = await sendEntryMessage(message)

        await sleep(1)
        await sendCountDownMessage(entryMessage, 60)

        await sleep(1)
        await sendCountDownMessage(entryMessage, 30)

        await sleep(1)
        await sendCountDownMessage(entryMessage, 10)

        await sleep(1)
        // メッセージを送付
        await sendStartMessage(entryMessage)
        console.log(entryMessage)
    }
});

// エントリーメッセージを送信します
const sendEntryMessage = async (message) => {
    const exampleEmbed = {
        color: 0x0099ff,
        title: '⚔️ Giveaway Battle ⚔️',
        description: '\nねだるな！勝ち取れ🔥🔥🔥\n',
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
                value: 'メッセージ送信から2分後()',
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

// 開始までのカウントダウンのメッセージを送信します
const sendCountDownMessage = async (entryMessage, second) => {
    const guildId = entryMessage.guildId
    const channelID = entryMessage.channelId
    const messageID = entryMessage.id

    const exampleEmbed = {
        color: 0x0099ff,
        title: `⚔️ Giveaway Battle開始まであと ${second}秒 ⚔️`,
        description: '\nAre You Ready?\n',
        fields: [
            {
                name: '参加リンク',
                // value: `${originalMessage.message}`,
                value: `[エントリーはこちら](https://discord.com/channels/${guildId}/${channelID}/${messageID})`,
            },
        ],
    };

    await entryMessage.channel.send({embeds: [exampleEmbed]})
}

// 開始メッセージを送信します
const sendStartMessage = async (entryMessage) => {
    // リアクションした人を取得します
    await getReactedUsers(entryMessage)

    // リアクションした人を文字列で結合します
    const entry = ""

    const exampleEmbed = {
        color: 0x0099ff,
        title: '⚔️ Battle Start❗️ ⚔️',
        description: entry,
        fields: [
            {
                name: '主催者',
                value: `${entryMessage.author}`,
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
                value: 'メッセージ送信から2分後()',
                inline: false,
            },
        ],
        timestamp: new Date(),
    };

    const sentMessage = await entryMessage.channel.send({embeds: [exampleEmbed]})

    try {
        await sentMessage.react("⚔️");
    } catch (error) {
        console.error("emoji failed to react")
    }

    return sentMessage;
}

// リアクションした人を集計します
const getReactedUsers = async (entryMessage) => {
    // const msg = await entryMessage.fetch({before: entryMessage.id, limit: 1})
    // console.log(await msg.reactions.cache.get("⚔️").users.cache)
    // const reactedUsers = [];
    //
    // return reactedUsers;
}

const sleep = (seconds) => new Promise(r => setTimeout(r, seconds * 1000));

client.login(process.env.APP_BOT_TOKEN);