import { useLanguageStore } from '../stores/language'

export function useI18n() {
  const languageStore = useLanguageStore()

  return {
    t: (key) => languageStore.t(key),
    currentLanguage: () => languageStore.currentLanguage,
    isRTL: () => languageStore.isRTL,
  }
}
