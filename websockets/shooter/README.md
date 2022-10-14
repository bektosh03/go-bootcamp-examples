# Task
**Online shooter**

Juda ham oddiy shooter yaratish kerak.
Shooter da 2 player bir biriga qarshi o'ynaydi.
Har birida 'shoot' (ya'ni otish) tugmasi (button) bo'ladi.
Har safar shu button bosilganda, qarshi playerni joni (100%)
10% ga kamayib boradi.
Kim birinchi raqibini jonini tugatsa, o'sha g'alaba qozongan bo'ladi.

Client tarafdan (ya'ni 'front' dan) otish tugmasi bosilganligi haqida
ma'lumot (qaysi player bosdi va qaysi playerga qarshi) yuborib turiladi
front qismi haqida o'ylash xozircha shartmas.

## NOTE
Qiziq tarafi:
- Otish tugmasi har safar random har xil joyda chiqadi.

## App Flow

1. Registratsiya qismi: bir formada ismini qabul qilsa bo'ldi (bu qismi fronti haqida ham o'ylash shartmas)
2. Matchmaking qismi: registratsiyadan o'tgan player lar ro'yxati dan o'ziga raqibi tanlaydi. NOTE: raqib 'accept' ya'ni o'ynashga rozi bo'lishi shartmas
3. Va asosiy tepada aytib o'tilgan o'yin jarayoni

## Development

**Concepts**:
- `command` - bu client tarafdan websocket orqali serverga keladigan xabar
- `event` - bu server tarafdan websocket orqali client ga yuboriladigan xabar

2-chi screen da ko'rsatilgan `Wait for someone` tugmasi bosilganda serverga quyidai xabar boradi:
```json
{
    "command": "wait_for_opponent",
    "player": { // bu shu tugmani bosgan playerning ma'lumomtlari
        "name": "<Player ismi>"
    }
}
```

2-chi screen da ko'rsatilgan `Available players` ro'xatida playerlar bo'lsa va ularning biri tanlansa, serverga quyidagi xabar boradi:
```json
{
    "command": "play",
    "player": { // bu shu tugmani bosgan playerning ma'lumomtlari
        "name": "<Player ismi>"
    },
    "rival": { // raqib player ma'lumotlari
        "name": "<Player ismi>"
    }
}
```
Va bu xabar qabul qilingandan keyin, raqibga (kutib turgan playerga) quyidagi xabar borishi kerak:
```json
{
    "event": "player_joined",
    "player": { // qo'shilgan player ma'lumotlari (ya'ni yuqoridagi player)
        "name": "<Player ismi>"
    }
}
```

Oxirgi etapda, xar bir player `Start` tugmasini bosishi kerak, u bosilganda serverga quyidagi xabar boradi:
```json
{
    "command": "start",
    "player": { // tugmani bosgan player
        "name": "<Player ismi>"
    }
}
```

Bu komanda qabul qilinganda server tekshirish kerak bo'ladi 2-chi player ham `Start` tugmasini bosganini, hali bosmagan bo'lsa, o'yin boshlanmay turadi,
ikkisi ham bosgandan keyin o'yin boshlanadi