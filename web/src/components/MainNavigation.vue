<template>
<div>
    <v-list nav>
        <v-list-item style="margin-top: 16px">
            <v-img src="data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiIHN0YW5kYWxvbmU9Im5vIj8+CjxzdmcKICAgeG1sbnM6ZGM9Imh0dHA6Ly9wdXJsLm9yZy9kYy9lbGVtZW50cy8xLjEvIgogICB4bWxuczpjYz0iaHR0cDovL2NyZWF0aXZlY29tbW9ucy5vcmcvbnMjIgogICB4bWxuczpyZGY9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkvMDIvMjItcmRmLXN5bnRheC1ucyMiCiAgIHhtbG5zOnN2Zz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciCiAgIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyIKICAgd2lkdGg9IjEzNC43NTc4M21tIgogICBoZWlnaHQ9IjEyOS42Nzk5OG1tIgogICB2aWV3Qm94PSIwIDAgMTM0Ljc1NzgzIDEyOS42Nzk5OCIKICAgdmVyc2lvbj0iMS4xIgogICBpZD0ic3ZnOCI+CjxkZWZzCiAgIGlkPSJkZWZzMiIgLz4KPG1ldGFkYXRhCiAgIGlkPSJtZXRhZGF0YTUiPgogICA8cmRmOlJERj4KICAgICAgPGNjOldvcmsKICAgICAgICAgcmRmOmFib3V0PSIiPgogICAgICA8ZGM6Zm9ybWF0PmltYWdlL3N2Zyt4bWw8L2RjOmZvcm1hdD4KICAgICAgPGRjOnR5cGUKICAgICAgICAgcmRmOnJlc291cmNlPSJodHRwOi8vcHVybC5vcmcvZGMvZGNtaXR5cGUvU3RpbGxJbWFnZSIgLz4KICAgICAgPGRjOnRpdGxlPjwvZGM6dGl0bGU+CiAgICAgIDwvY2M6V29yaz4KICAgPC9yZGY6UkRGPgo8L21ldGFkYXRhPgo8ZwogICBpZD0ibGF5ZXIxIgogICB0cmFuc2Zvcm09InRyYW5zbGF0ZSgtNTMuOTk4MDc1LC02Ny45MjIyNjYpIj4KICAgPHBhdGgKICAgICAgc3R5bGU9ImZpbGw6IzRjNzMzNztmaWxsLW9wYWNpdHk6MTtzdHJva2Utd2lkdGg6MC4yNjQ1ODMiCiAgICAgIGQ9Im0gNTYuNTE2MjY1LDE5Ny4yNjQwMiBjIC0xLjczOTE5LC0wLjQ2NzQyIC0yLjMzNTYzLC0xLjEyNzAxIC0yLjQ5MjE2LC0yLjc1NjA4IC0wLjI1MzcxLC0yLjY0MDM3IDEuMzA2OTksLTMuODg4MTcgNi40NTY5MSwtNS4xNjIzOSA3LjQ1ODI4LC0xLjg0NTM3IDEzLjAzMDEzLC01LjM4OTY3IDE3LjExMDE1LC0xMC44ODM5IDIuODQzNzQsLTMuODI5NDIgNy41MDAyOSwtMTQuMjE1MDggMTMuMDM3NjksLTI5LjA3ODQgbCAyLjM1NjEzLC02LjMyNDIzIDEuNzIwNiwtMC43NzgwMiBjIDIuNTk1ODksLTEuMTczODEgMTMuNjg0Njk1LC01Ljc2ODQ4IDIwLjE1NzQwNSwtOC4zNTIyNyAzLjE1NTIzLC0xLjI1OTUxIDUuNjk0MTIsLTIuMzMyNjkgNS42NDE5NiwtMi4zODQ4NSAtMC4wNTIyLC0wLjA1MjIgLTMuMjI5MDMsMC44MDAyIC03LjA1OTcxLDEuODk0MTMgLTEwLjkyMDU5LDMuMTE4NjEgLTE4LjA2NzgzNSw0LjkwNDcyIC0xOC4zNzk1MDUsNC41OTMwNSAtMC4yNDc3NCwtMC4yNDc3NCAzLjE2MzcxLC04LjMxODE1IDUuODUxNjA1LC0xMy44NDMwNSBsIDIuNzE5NzIsLTUuNTkwMzEgMS4yODkzNiwtMC40MjQwMyBjIDEuMzA0NjYsLTAuNDI5MDcgOC45MDgyNCwtMi4yOTA4NCAxMy42NzY4NSwtMy4zNDg4MyAxLjQyODg1LC0wLjMxNzAyIDIuNTU0MTEsLTAuNjIwMiAyLjUwMDU3LC0wLjY3Mzc0IC0wLjA1MzUsLTAuMDUzNSAtMy40NzExNCwwLjE0MDQyIC03LjU5NDY3LDAuNDMxMDIgLTQuMTIzNTMsMC4yOTA2MSAtNy41MzkzOCwwLjQ4NjMyIC03LjU5MDc5LDAuNDM0OTIgLTAuMTM4MjEsLTAuMTM4MjIgMy4yODU0OSwtNS4xMjM4IDUuMzUwNDgsLTcuNzkxMzUgNC44OTYyMywtNi4zMjQ5MiAxMS4zNjk0NywtMTEuNzg2Njg3IDE3LjgzMjg0LC0xNS4wNDYzNjcgMS4zMDk5OCwtMC42NjA2NjUgMi4zODE3OCwtMS4yNDgwMTIgMi4zODE3OCwtMS4zMDUyMTYgMCwtMC4yNDAwMzcgLTUuNTg2MjIsMS44NTk5MTMgLTcuNDc5MzQsMi44MTE2MDYgLTQuMzgwNDIsMi4yMDIwODcgLTkuODk5ODMsNi4wNTcyNTUgLTEzLjMyNTY4LDkuMzA3NjY3IC0wLjc0MjIzLDAuNzA0MjIgLTEuMTU1NjEsMC45MzcxNyAtMS4yODY2OSwwLjcyNTA5IC0wLjEwNjIxLC0wLjE3MTg2IC0wLjYzNjEzLC0yLjUzNDQ4IC0xLjE3NzYsLTUuMjUwMjcxIGwgLTAuOTg0NDksLTQuOTM3Nzk5IC0wLjE0OTUsNi41MzczNCAtMC4xNDk1MSw2LjUzNzM0IC0xLjYwMzc5LDEuNzk3MDMgYyAtMy40NDY3NywzLjg2MjA5IC03LjU2NDQ4NSw5LjYyOTM1IC0xMC40OTEzNTUsMTQuNjk0MTggLTAuNTY3MjIsMC45ODE1NSAtMS4xMTEzNywxLjY5NTkzIC0xLjIwOTIzLDEuNTg3NSAtMC4wOTc5LC0wLjEwODQzIC0yLjE2NjQzLC00LjA5NjQ0IC00LjU5NjgzLC04Ljg2MjI0IC0yLjQzMDQxLC00Ljc2NTgxIC00LjQ1ODE2LC04LjY2NTExIC00LjUwNjEyLC04LjY2NTExIC0wLjEwMDkyLDAgMC4xMDYyOSwwLjY3OTc4IDMuMDEyNjUsOS44ODI5IDEuMTg4MDgsMy43NjIxMSAyLjUxNDM0LDguMTAxNjIgMi45NDcyNiw5LjY0MzM2IGwgMC43ODcxMSwyLjgwMzE3IC0xLjg4MTUxLDMuODUwMzkgYyAtMi4zMDc2NCw0LjcyMjQyIC01LjAzNDI3LDExLjA2MTg2IC03LjkxNTAyLDE4LjQwMjQ3IC0xLjE5OTI3LDMuMDU1OTQgLTIuMjU3MTksNS42NDIxMiAtMi4zNTA5Miw1Ljc0NzA2IC0wLjE5OTA4LDAuMjIyODggLTMuMTg1NTYsLTMuNDE0NDcgLTQuOTYyMjEsLTYuMDQzNjcgLTcuMDczNjEsLTEwLjQ2Nzk0IC04LjYxMjkyLC0yNS45NzQ3NiAtMy45NTkwMiwtMzkuODgyNzIgNC42NTk2MywtMTMuOTI1MTAxIDE1LjQxMjU2LC0yNS40NjUwNDggMzAuNzIyNDI1LC0zMi45NzEwMzYgMTEuODg3MDYsLTUuODI3ODgzIDI0LjA2NTcsLTkuMDA2MDM2IDM5LjY4NzUsLTEwLjM1Njg5OCA0Ljc3Mzg2LC0wLjQxMjgxIDE2LjkzMzk4LC0wLjQxMjEyMiAyMS42OTU4MywwLjAwMTIgOS42NDYwMiwwLjgzNzMxMyAxOS45NDc0NSwyLjM3MDUwOSAyNS4zOTUwNywzLjc3OTYzIGwgMS4wNTM0LDAuMjcyNDggLTAuNzg4ODIsMC40MTEyIGMgLTAuNDMzODUsMC4yMjYxNTkgLTIuMDk4NSwxLjQyNDQ4OSAtMy42OTkyMywyLjY2Mjk1NSAtMTAuNzMwMTYsOC4zMDE3OTUgLTE2LjE4MjQyLDE0Ljc5NzU2OCAtMjAuMTM4MDQsMjMuOTkyMjY3IC0yLjA3ODc5LDQuODMyMDkyIC0yLjgxNDMsNy4yNzIyNjIgLTUuNjYxMiwxOC43ODE5MDIgLTQuMDkzMzIsMTYuNTQ4NzcgLTUuMTQxMDgsMTkuODg2IC04LjMwODA0LDI2LjQ2MTg1IC0zLjI1MTYzLDYuNzUxNjggLTYuODk4NzMsMTEuOTA4IC0xMS42Nzk1OCwxNi41MTI3OSAtNy4xMTQ4Miw2Ljg1MjggLTE0Ljk4NjQ1LDEwLjA5MjM3IC0yNC41MjQyNCwxMC4wOTI5NyAtNy40MDExNCw1LjNlLTQgLTEzLjE2MjgsLTEuMTc5NDIgLTIwLjY2NDIxNSwtNC4yMzE2NCAtMS43NzgyNywtMC43MjM1NSAtMi42MTUzMiwtMC45NTEyMiAtMi42OTExNiwtMC43MzE5NSAtMC4wNjA1LDAuMTc1MDMgLTAuNTQwNjcsMS42Mjc5NCAtMS4wNjY5NywzLjIyODY2IC01LjAwMzU3LDE1LjIxODM2IC0xMy43NDAzNiwyNC4zMDAzMiAtMjYuNDg3MzksMjcuNTMzODcgLTIuODA1MiwwLjcxMTYgLTQuNTk2NzIsMC43ODQxMSAtNi41MzA3NiwwLjI2NDM0IHoiCiAgICAgIGlkPSJwYXRoODU3IiAvPgo8L2c+Cjwvc3ZnPgo=" :aspect-ratio="1/1" max-height="64" max-width="64"></v-img>
        </v-list-item>

        <v-list-item style="margin-top: -32px">
            <v-list-item-content>
                <v-list-item-title class="title">
                    Paperless Office
                </v-list-item-title>
                <v-list-item-subtitle>Rethink sustainability</v-list-item-subtitle>
            </v-list-item-content>
        </v-list-item>
    </v-list>
    <v-divider class="mx-4"></v-divider>
    <v-list nav shaped v-for="(navSection, sectionTitle) in value" :key="sectionTitle">
        <div class="my-3"></div>
        <v-subheader class="text-uppercase" v-text="sectionTitle"></v-subheader>
        <v-list-item-group color="dark">
            <v-list-item v-for="(navItem, index) in navSection" :key="index" :to="navItem.to">
                <v-list-item-icon>
                    <v-icon v-text="navItem.icon"></v-icon>
                </v-list-item-icon>
                <v-list-item-content>
                    <v-badge v-if="navItem.title == 'Unconfirmed' && unconfirmedNum > 0" color="secondary" :content="unconfirmedNum" inline>
                        <v-list-item-title v-text="navItem.title"></v-list-item-title>
                    </v-badge>
                    <v-list-item-title v-else v-text="navItem.title"></v-list-item-title>
                </v-list-item-content>
            </v-list-item>
        </v-list-item-group>
    </v-list>
</div>
</template>

<script>
export default {
    props: ["value", "unconfirmedNum"]
}
</script>
