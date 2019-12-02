import discord
import os
import sys

guildid = os.getenv("GUILDID", None)
token = os.getenv("TOKEN", None)

if not (guildid or token):
    print("STUFF IS MISSING")
    sys.exit(0)

print(guildid)
client = discord.Client()

@client.event
async def on_ready():
    print("READY")
    print(client.user.name)

    guild = client.get_guild(int(guildid))

    for role in guild.roles:
        print("----------------")
        print(role.name)
        print(role.id)

client.run(token)
