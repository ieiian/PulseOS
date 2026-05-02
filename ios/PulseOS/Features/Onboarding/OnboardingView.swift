import SwiftUI

struct OnboardingView: View {
    @State var profile: UserProfile
    let onComplete: (UserProfile) -> Void

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 16) {
                Text("欢迎来到 PulseOS")
                    .font(.title2.weight(.semibold))
                Text("先完成基础资料和健康目标设置。")
                    .foregroundStyle(.secondary)

                Group {
                    TextField("姓名", text: $profile.name)
                    TextField("年龄", text: $profile.age)
                    TextField("性别", text: $profile.gender)
                    TextField("身高（cm）", text: $profile.heightCM)
                    TextField("体重（kg）", text: $profile.weightKG)
                }
                .textFieldStyle(.roundedBorder)

                GoalPickerView(selected: $profile.primaryGoal)

                Button("完成设置") {
                    onComplete(profile)
                }
                .buttonStyle(.borderedProminent)
                .frame(maxWidth: .infinity, alignment: .center)
            }
            .padding(24)
        }
    }
}

private struct GoalPickerView: View {
    @Binding var selected: String

    private let goals = [
        ("fat_loss", "减脂"),
        ("maintain", "维持健康"),
        ("muscle_gain", "增肌"),
        ("sleep_recovery", "作息恢复"),
        ("stress_relief", "减压放松"),
    ]

    var body: some View {
        VStack(alignment: .leading, spacing: 8) {
            Text("健康目标")
                .font(.headline)
            LazyVGrid(columns: [GridItem(.adaptive(minimum: 110))], spacing: 8) {
                ForEach(goals, id: \.0) { goal in
                    Button(goal.1) {
                        selected = goal.0
                    }
                    .buttonStyle(.borderedProminent)
                    .tint(selected == goal.0 ? PulseTheme.brand : .gray.opacity(0.35))
                }
            }
        }
    }
}

