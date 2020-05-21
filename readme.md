<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [playstore-ratings](#playstore-ratings)
- [how to run](#how-to-run)
  - [example](#example)
- [how to build](#how-to-build)
  - [for your current platform](#for-your-current-platform)
  - [for another platform](#for-another-platform)
  - [for all platforms](#for-all-platforms)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# playstore-ratings
this application pulls the playstore ratings for all applications on a landing page

# how to run
by executing the binary it will pull the ratings of all the applications for a landing page and populate your copy/paste buffer with the results.  You can then 'paste' those results into Excel.

double click on the binary for default behavior, which pulls playstore ratings for verizon apps -> [default url used](https://play.google.com/store/apps/collection/cluster?clp=igM6ChkKEzgyMDQ2OTkzNjYyNDAwMTk3MDQQCBgDEhsKFWNvbS52encuaHNzLm15dmVyaXpvbhABGAMYAQ%3D%3D:S:ANO1ljK_y6A&gsr=Cj2KAzoKGQoTODIwNDY5OTM2NjI0MDAxOTcwNBAIGAMSGwoVY29tLnZ6dy5oc3MubXl2ZXJpem9uEAEYAxgB:S:ANO1ljLQ2zk&gl=US)

or run it from the command line and pass the landing page url that you'd like to pull ratings for

## example
to execute from command line for Mac:

```
bin/playstore-ratings-darwin https://play.google.com/store/apps
```

results in the below (tab delimited for [Excel](https://www.microsoft.com/en-ie/microsoft-365/excel) pasting consumption) being copied into your copy/paste buffer:
```
App Name	App Total	App Rating	App URL
Yu-Gi-Oh! Duel Links	1,949,483	4.5	https://play.google.com/store/apps/details?id=jp.konami.duellinks
KartRider Rush+	60,754	3.9	https://play.google.com/store/apps/details?id=com.nexon.kart
Board Kings™️	443,733	4.3	https://play.google.com/store/apps/details?id=com.jellybtn.boardkings
Mining Inc.	471	4.2	https://play.google.com/store/apps/details?id=com.gpp.mininginc
Sniper 3D: Fun Offline Gun Shooting Games Free	12,893,923	4.4	https://play.google.com/store/apps/details?id=com.fungames.sniper3d
Soccer Stars	2,114,267	4.2	https://play.google.com/store/apps/details?id=com.miniclip.soccerstars
D.B.System	3,797	4.6	https://play.google.com/store/apps/details?id=com.Company.BreathofDragons
Crafty Candy Blast	4,075	4.6	https://play.google.com/store/apps/details?id=com.outplayentertainment.craftycandyblast
Disney Getaway Blast	40,233	4.6	https://play.google.com/store/apps/details?id=com.gameloft.anmp.disney.adventure
Piggy GO - Clash of Coin	181,038	4.6	https://play.google.com/store/apps/details?id=com.aladinfun.piggytravel.android
Messenger – Text and Video Chat for Free	72,526,789	4.2	https://play.google.com/store/apps/details?id=com.facebook.orca
Facebook	99,366,667	4.3	https://play.google.com/store/apps/details?id=com.facebook.katana
Facebook Lite	14,259,617	4.2	https://play.google.com/store/apps/details?id=com.facebook.lite
Instagram	99,599,750	4.5	https://play.google.com/store/apps/details?id=com.instagram.android
Snapchat	22,220,540	4.4	https://play.google.com/store/apps/details?id=com.snapchat.android
TikTok - Make Your Day	28,236,135	1.3	https://play.google.com/store/apps/details?id=com.zhiliaoapp.musically
Viber Messenger - Messages, Group Chats & Calls	13,765,843	4.3	https://play.google.com/store/apps/details?id=com.viber.voip
Spotify: Listen to new music, podcasts, and songs	18,745,845	4.6	https://play.google.com/store/apps/details?id=com.spotify.music
Subway Surfers	32,789,570	4.5	https://play.google.com/store/apps/details?id=com.kiloo.subwaysurf
8 Ball Pool	18,892,838	4.4	https://play.google.com/store/apps/details?id=com.miniclip.eightballpool
WhatsApp Messenger	112,328,074	4.3	https://play.google.com/store/apps/details?id=com.whatsapp
Snapchat	22,220,650	4.4	https://play.google.com/store/apps/details?id=com.snapchat.android
Spotify: Listen to new music, podcasts, and songs	18,745,551	4.6	https://play.google.com/store/apps/details?id=com.spotify.music
TikTok - Make Your Day	28,204,707	1.3	https://play.google.com/store/apps/details?id=com.zhiliaoapp.musically
Instagram	99,600,376	4.5	https://play.google.com/store/apps/details?id=com.instagram.android
Facebook	99,366,070	4.3	https://play.google.com/store/apps/details?id=com.facebook.katana
Adverts.ie Buy & Sell Nearby	21,269	4.5	https://play.google.com/store/apps/details?id=com.distilledmedia.adverts
DoneDeal - New & Used Cars For Sale	19,341	4.4	https://play.google.com/store/apps/details?id=ie.donedeal.android
AIB Mobile	10,980	3.5	https://play.google.com/store/apps/details?id=aib.ibank.android
Roblox	13,270,222	4.4	https://play.google.com/store/apps/details?id=com.roblox.client
Subway Surfers	32,789,510	4.5	https://play.google.com/store/apps/details?id=com.kiloo.subwaysurf
Clash Royale	28,634,604	4.3	https://play.google.com/store/apps/details?id=com.supercell.clashroyale
Parchisi STAR Online	1,306,953	4.5	https://play.google.com/store/apps/details?id=com.superking.parchisi.star
Tiles Hop: EDM Rush!	1,152,894	4.2	https://play.google.com/store/apps/details?id=com.amanotes.beathopper
Save The Girl	371,559	3.5	https://play.google.com/store/apps/details?id=com.xmgame.savethegirl
Video Editor & Video Maker - InShot	5,828,490	4.8	https://play.google.com/store/apps/details?id=com.camerasideas.instashot
PicsArt Photo Editor: Pic, Video & Collage Maker	9,467,144	4.3	https://play.google.com/store/apps/details?id=com.picsart.studio
BeautyPlus - Easy Photo Editor & Selfie Camera	4,433,823	4.4	https://play.google.com/store/apps/details?id=com.commsource.beautyplus
B612 - Beauty & Filter Camera	6,549,433	4.3	https://play.google.com/store/apps/details?id=com.linecorp.b612.android
Ulike - Define your selfie in trendy style	388,443	4.6	https://play.google.com/store/apps/details?id=com.gorgeous.liteinternational
Home Workout for Women - Female Fitness	15,265	4.8	https://play.google.com/store/apps/details?id=workoutforwomen.femalefitness.womenworkout.loseweight
Co–Star Personalized Astrology	14,920	4.7	https://play.google.com/store/apps/details?id=com.costarastrology
GroupCal - Shared Calendar	349	4.3	https://play.google.com/store/apps/details?id=app.groupcal.www
PDF Editor - All-powerful PDF Reader & Manager	7,010	3.0	https://play.google.com/store/apps/details?id=com.kmo.pdf.editor
byte - short looping videos	2,837	3.8	https://play.google.com/store/apps/details?id=co.byte
Calm Colors - Coloring Book	193	4.1	https://play.google.com/store/apps/details?id=com.vividgames.calm.colors.coloring.book
Canva: Graphic Design, Video, Invite & Logo Maker	2,601,816	4.7	https://play.google.com/store/apps/details?id=com.canva.editor
FitOn - Free Fitness Workouts & Personalized Plans	13,330	4.8	https://play.google.com/store/apps/details?id=com.fiton.android
Color By Number For Adults	56,924	4.6	https://play.google.com/store/apps/details?id=com.pixign.premium.coloring.book
Wallet - Money, Budget, Finance & Expense Tracker	143,255	4.6	https://play.google.com/store/apps/details?id=com.droid4you.application.wallet
Video Editor - Glitch Video Effects	170,114	4.6	https://play.google.com/store/apps/details?id=glitchvideoeditor.videoeffects.glitchvideoeffect
Disney Collect! by Topps	3,235	4.1	https://play.google.com/store/apps/details?id=com.topps.disney
Babbel - Learn Languages - Spanish, French & More	509,597	4.6	https://play.google.com/store/apps/details?id=com.babbel.mobile.android.en
Sleep Sounds	102,025	4.7	https://play.google.com/store/apps/details?id=net.metapps.sleepsounds
Candy Camera - selfie, beauty camera, photo editor	3,495,156	4.4	https://play.google.com/store/apps/details?id=com.joeware.android.gpulumera
Any.do: To do list, Calendar, Planner & Reminders	346,324	4.4	https://play.google.com/store/apps/details?id=com.anydo
green	4,723	4.6	https://play.google.com/store/apps/details?id=air.com.bartbonte.green
Dungeon Corporation VIP: An auto-farming RPG game!	235	4.0	https://play.google.com/store/apps/details?id=com.bigshotgames.legendaryItemVip
Fancade	1,476	4.2	https://play.google.com/store/apps/details?id=com.martinmagni.fancade
Rocky Rampage: Wreck 'em Up	2,504	4.2	https://play.google.com/store/apps/details?id=com.joyseed.bod
Raccoon Journey: Match-3 Puzzle Adventure 2020	35	4.1	https://play.google.com/store/apps/details?id=mobi.blackbears.raccoon_journey
Slash & Girl - Joker World	4,121	4.0	https://play.google.com/store/apps/details?id=com.slash.girl.redfish
Life Gallery	2,693	3.9	https://play.google.com/store/apps/details?id=com.Games751.LifeGallery
Door Kickers: Action Squad	316	4.6	https://play.google.com/store/apps/details?id=com.khg.actionsquad
Bomb Chicken	34	4.6	https://play.google.com/store/apps/details?id=com.nitrome.bombchicken
Opera Mini - fast web browser	6,884,276	4.4	https://play.google.com/store/apps/details?id=com.opera.mini.native
Yandex	805,554	4.3	https://play.google.com/store/apps/details?id=ru.yandex.searchplugin
Tor Browser	34,295	4.2	https://play.google.com/store/apps/details?id=org.torproject.torbrowser
Puffin Web Browser	728,916	3.9	https://play.google.com/store/apps/details?id=com.cloudmosa.puffinFree
```


# how to build

## for your current platform
```
make
```

## for another platform
```
make linux
make windows
make darwin
```

## for all platforms
```
make all
```
