import SwiftUI

struct UserView: View {
    let profile: UserProfile

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 16) {
                Text("用户")
                    .font(.title2.weight(.semibold))
                UserSection(title: "个人信息", body: "\(profile.name.isEmpty ? "未设置姓名" : profile.name) · \(profile.age.isEmpty ? "-" : profile.age) 岁")
                UserSection(title: "健康目标", body: goalLabel(profile.primaryGoal))
                UserSection(title: "身体指标", body: "\(profile.heightCM.isEmpty ? "-" : profile.heightCM) cm / \(profile.weightKG.isEmpty ? "-" : profile.weightKG) kg")
                UserSection(title: "统计数据", body: "连续记录 0 天，后续接后端统计接口")
                Divider()
                UserSection(title: "设置", body: "通知、权限、设备管理结构已预留")
            }
            .padding(16)
        }
        .background(PulseTheme.background)
    }

    private func goalLabel(_ goal: String) -> String {
        switch goal {
        case "fat_loss":
            return "减脂"
        case "muscle_gain":
            return "增肌"
        case "sleep_recovery":
            return "作息恢复"
        case "stress_relief":
            return "减压放松"
        default:
            return "维持健康"
        }
    }
}

private struct UserSection: View {
    let title: String
    let body: String

    var body: some View {
        VStack(alignment: .leading, spacing: 6) {
            Text(title)
                .font(.headline)
            Text(body)
                .foregroundStyle(.secondary)
        }
    }
}

