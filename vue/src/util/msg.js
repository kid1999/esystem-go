import { Notify } from 'quasar'

export const successMsg = (msg) => {
  Notify.create({ color: 'teal', message: msg, icon: 'tag_faces', position: 'top' })
}

export const failMsg = (msg) => {
  Notify.create({ message: msg, icon: 'warning', position: 'top' })
}

export const timeoutMsg = (msg) => {
  Notify.create({ message: msg, icon: 'warning', position: 'top', spinner: true, timeout: 2000 })
}
