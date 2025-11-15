const imageToCanvas = async (imageFile: File): Promise<HTMLCanvasElement> => {
  return new Promise((resolve, _reject) => {
    const reader = new FileReader()
    reader.onload = (ev) => {
      const image = new Image()

      image.onload = () => {
        const canvas = document.createElement("canvas")
        const ctx = canvas.getContext("2d")

        if (!ctx) return

        const targetWidth = 800
        const ratio = targetWidth / image.width
        const width = image.width * ratio
        const height = image.height * ratio

        canvas.width = width
        canvas.height = height

        ctx.drawImage(image, 0, 0, width, height)
        resolve(canvas)
      }

      if (!ev.target || !ev.target.result) return

      image.src = ev.target.result as string
    }

    reader.readAsDataURL(imageFile)
  })
}

export const imageToImageURL = async (imageFile: File): Promise<string> => {
  return (await imageToCanvas(imageFile)).toDataURL()
}

export const imageToBlob = async (imageFile: File): Promise<Blob | null> => {
  return new Promise(async (resolve, _reject) => {
    ;(await imageToCanvas(imageFile)).toBlob(resolve, "image/jpeg", 0.7)
  })
}
