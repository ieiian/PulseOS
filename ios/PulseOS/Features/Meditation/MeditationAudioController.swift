import Foundation

final class MeditationAudioController: ObservableObject {
    @Published private(set) var isPlaying = false

    func toggle() {
        isPlaying.toggle()
    }
}

